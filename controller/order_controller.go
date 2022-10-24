package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
	"webProject/service"
	"webProject/util"
)

type OrderController struct {
	Session *sessions.Session
	Server  service.OrderDetailServer
	Ctx     iris.Context
}

/*
*get请求
*请求路径/bos/orders
*请求参数offset、limit
 */

func (oc OrderController) Get() mvc.Response {
	limitStr := oc.Ctx.FormValue("limit")
	offsetStr := oc.Ctx.FormValue("offset")

	if limitStr == "" || offsetStr == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  util.RECODE_FAIL,
				"type":    util.RESPMSG_ERROR_ORDERLIST,
				"massage": util.Recode2Text(util.RESPMSG_ERROR_ORDERLIST),
			},
		}
	}
	limit, err := strconv.Atoi(limitStr)
	offset, err := strconv.Atoi(offsetStr)

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  util.RECODE_FAIL,
				"type":    util.RESPMSG_ERROR_ORDERLIST,
				"massage": util.Recode2Text(util.RESPMSG_ERROR_ORDERLIST),
			},
		}
	}
	if limit > 50 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	orderList := oc.Server.GetOrderList(offset, limit)

	//对返回结果按照前端进行组装
	var respList []interface{}

	for _, order := range orderList {
		respList = append(respList, order.OrderDetail2Resp())
	}
	return mvc.Response{
		Object: &respList,
	}
}
