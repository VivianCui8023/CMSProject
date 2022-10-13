package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"webProject/service"
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
