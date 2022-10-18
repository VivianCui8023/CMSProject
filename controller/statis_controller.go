package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strings"
	"webProject/service"
	"webProject/util"
)

// 统计信息处理
type StatisController struct {
	Ctx     iris.Context
	Service service.StatisServer

	Session *sessions.Session
}

// 路径信息/statis/order/NaN-NaN-NaN/count

func (sc *StatisController) GetCount() mvc.Response {
	path := sc.Ctx.Path()
	var pathSlice []string
	if path != "" {
		//直接分割时候"///"路径前会有空
		pathSlice = strings.Split(path, "/")
	}
	//去掉空的
	pathSlice = pathSlice[1:]
	fmt.Printf("len(pathSlice)%d\n", len(pathSlice))
	if len(pathSlice) != 4 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": util.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	mode := pathSlice[1]
	getDate := pathSlice[2]
	fmt.Printf("mode%s,getdate%s", mode, getDate)
	var result int64
	switch mode {
	case "admin":
		result = sc.Service.GetAdminDailyCount(getDate)
	case "user":
		result = sc.Service.GetUserDailyCount(getDate)
	case "order":
		result = sc.Service.GetOrderDailyCount(getDate)

	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": util.RECODE_OK,
			"count":  result,
		},
	}
}
