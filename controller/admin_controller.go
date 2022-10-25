package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"io"
	"os"
	"strconv"
	"webProject/model"
	"webProject/service"
	"webProject/util"
)

/*我们使用mvc包模式来进行功能开发，在进行了结构体定义以后，我们接着定义控制器。
//控制器负责来完成我们请求的逻辑流程控制，是我们功能开发的核心枢纽。
//在AdminController定义中，包含iris.Context上下文处理对象，
//用于数据功能处理的管理员模块功能实现AdminService，
//还有用于session管理的对象。定义PostLogin方法来处理用户登陆请求。*/

type AdminController struct {
	//上下文处理
	Ctx iris.Context
	//服务
	Service service.AdminService
	//session
	Session *sessions.Session
}

const ADMIN = "admin" //用于存储在Session中

// 存储post中用户密码数据结构,以及对应json转义
type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//真正的接口实现在这,利用iris的mvc模式中自动寻找处理
/**
* 管理员登录功能：json请求格式
* 路径：/admin/login
* 请求方式:post
* 返回参数{
		"status":  "0",
		"success": "登录失败",
		"message": "用户名或密码为空,请重新填写后尝试登录",
			}
*/

func (ac AdminController) PostLogin(context iris.Context) mvc.Response {
	//先从请求中获取用户输入name跟pwd，json格式

	var adminLogin AdminLogin
	//ac.Ctx = context
	err := ac.Ctx.ReadJSON(&adminLogin)
	if err != nil {
		panic(err.Error())
	}
	println("------adminLogin--------" + adminLogin.UserName + adminLogin.Password)

	//空参数检测
	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "用户名或密码为空",
			},
		}
	}

	//送到service查询
	admin, isExist := ac.Service.GetAdminNameAndPassword(adminLogin.UserName, adminLogin.Password)
	if !isExist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "用户名或密码错误",
			},
		}
	}

	//用户存在，将信息存到session中
	userByte, _ := json.Marshal(admin)
	fmt.Printf("session中信息:%s\n", string(userByte))
	//这个地方有空指针问题
	ac.Session.Set(ADMIN, userByte)

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  "1",
			"success": "登录成功",
			"message": "管理员登录成功",
		},
	}
}

// 在session中返回管理员信息 路径/admin/info

func (ac *AdminController) GetInfo() mvc.Response {
	userByte := ac.Session.Get(ADMIN)
	if userByte == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  util.RECODE_UNLOGIN,
				"type":    util.EEROR_UNLOGIN,
				"message": util.Recode2Text(util.EEROR_UNLOGIN),
			},
		}
	}

	var admin model.Admin
	err := json.Unmarshal(userByte.([]byte), &admin)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  util.RECODE_UNLOGIN,
				"type":    util.EEROR_UNLOGIN,
				"message": util.Recode2Text(util.EEROR_UNLOGIN),
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": util.RECODE_OK,
			"data":   admin.AdmintoRespone(),
		},
	}
}

// 退出登录
// 请求方式：get
// 请求路径/admin/singout
func (ac *AdminController) GetSingout() mvc.Response {
	//这里就是删除了Session中的东西
	ac.Session.Delete(ADMIN)
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  util.RECODE_OK,
			"success": util.Recode2Text(util.RESPMSG_SIGNOUT),
		},
	}
}

// 获取管理员数量接口路径/admin/count
func (ac *AdminController) GetCount() mvc.Response {
	total, err := ac.Service.GetAdminCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  util.RECODE_FAIL,
				"message": util.Recode2Text(util.RESPMSG_ERRORADMINCOUNT),
				"count":   0,
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": util.RECODE_OK,
			"count":  total,
		},
	}
}

type LimitOffset struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (ac *AdminController) PostInfo() mvc.Response {
	var para LimitOffset
	err := ac.Ctx.ReadJSON(&para)
	if err != nil {
		panic(err.Error())
	}
	info := make([]model.Admin, 0)
	info = ac.Service.PostAdminInfo(para.Limit, para.Offset)
	return mvc.Response{
		Object: map[string]interface{}{
			"status": util.RECODE_OK,
			"info":   info,
		},
	}
}

// 管理员用户头像上传/admin/update/avatar/1 post
// BeforeActivation支持自定义回调事件，用来自定义控制器与处理路径，其他的会自动处理匹配

func (ac AdminController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/update/avatar/{adminId}", "PostAvatar")
}

func (ac AdminController) PostAvatar() mvc.Response {
	adminID := ac.Ctx.Params().Get("adminId")
	iris.New().Logger().Info(adminID)
	file, info, err := ac.Ctx.FormFile("file")
	if err != nil {
		//这个地方不能用panic
		iris.New().Logger().Error(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  util.RECODE_FAIL,
				"type":    util.RESPMSG_ERROR_PICTUREADD,
				"message": util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD),
			},
		}

	}
	defer file.Close()
	fileName := info.Filename
	//文件放在本地目录upload
	out, err := os.OpenFile("./uploads/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		iris.New().Logger().Error(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  util.RECODE_FAIL,
				"type":    util.RESPMSG_ERROR_PICTUREADD,
				"message": util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD),
			},
		}
	}
	iris.New().Logger().Info("文件路径：" + out.Name())
	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
		iris.New().Logger().Error(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  util.RECODE_FAIL,
				"type":    util.RESPMSG_ERROR_PICTUREADD,
				"message": util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD),
			},
		}
	}
	intAdminId, _ := strconv.Atoi(adminID)
	//更新头像信息到数据库中
	err = ac.Service.PostAvatar(intAdminId, fileName)
	return mvc.Response{
		Object: map[string]interface{}{
			"status":   util.RECODE_OK,
			"img_path": fileName,
		},
	}

}
