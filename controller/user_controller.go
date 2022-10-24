package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
	"webProject/service"
	"webProject/util"
)

type UserController struct {
	Ctx     iris.Context
	Session *sessions.Session
	Service service.UserService
}

/**
 * 获取用户总数
 * 请求类型：Get
 * 请求Url：/v1/users/count
 */

func (uc *UserController) GetCount() mvc.Response {

	result, err := uc.Service.GetUserCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": util.RECODE_FAIL,
				"count":  0,
			},
		}
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"status": util.RECODE_OK,
			"count":  result,
		},
	}
}

/**
 * 获取用户总数
 * 请求类型：Get
 * 请求Url：/v1/users/list
 */

func (uc *UserController) GetList() mvc.Response {
	limitStr := uc.Ctx.FormValue("limit")
	offsetStr := uc.Ctx.FormValue("offset")

	limit, err := strconv.Atoi(limitStr)
	offset, err := strconv.Atoi(offsetStr)

	if limitStr == "" || offsetStr == "" || err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  util.RECODE_FAIL,
				"message": util.Recode2Text(util.RESPMSG_ERROR_USERLIST),
				"type":    util.RESPMSG_ERROR_USERLIST,
			},
		}
	}
	if offset < 0 {
		offset = 0
	}
	if limit > 50 {
		limit = 50
	}
	userList := uc.Service.GetUserList(offset, limit)

	if len(userList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  util.RECODE_FAIL,
				"message": util.Recode2Text(util.RESPMSG_ERROR_USERLIST),
				"type":    util.RESPMSG_ERROR_USERLIST,
			},
		}
	}
	fmt.Println("----------------查询情况----------------")
	fmt.Println(userList)
	fmt.Println("----------------返回效果----------------")
	//将查询到的信息按照前端进行组装
	var respList []interface{}
	for _, user := range userList {
		respList = append(respList, user.UsertoResponeDec())
	}
	fmt.Println(respList)

	return mvc.Response{
		Object: &respList,
	}

}
