package controller

import (
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"webProject/service"
	"webProject/util"
)

type UserController struct {
	Ctx     context.Context
	Session *sessions.Session
	Service service.UserService
}

func (uc *UserController) GetCount() mvc.Response {
	//是否判断用户登录？
	result := uc.Service.GetUserCount()
	return mvc.Response{
		Object: map[string]interface{}{
			"status": util.RECODE_OK,
			"count":  result,
		},
	}
}
