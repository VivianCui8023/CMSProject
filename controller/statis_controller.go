package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strings"
	"time"
	"webProject/service"
	"webProject/util"
)

// 统计信息处理

type StatisController struct {
	Ctx     iris.Context
	Service service.StatisServer

	Session *sessions.Session
}

const (
	ADMINMODEL = "adminModel"
	ORDERMODEL = "orderModel"
	USERMODEL  = "userModel"
)

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
	if getDate == "NaN-NaN-NaN" {
		getDate = time.Now().Format("2006-01-02")
	}
	var result int64
	switch mode {
	case "admin":
		sessionStatus := sc.Session.Get(ADMINMODEL + getDate)
		if sessionStatus == nil {
			result = sc.Service.GetAdminDailyCount(getDate)
			sc.Session.Set(ADMINMODEL+getDate, result)
		} else {
			sessionResult := sessionStatus.(int64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": util.RECODE_OK,
					"count":  sessionResult,
				},
			}
		}

	case "user":
		sessionStatus := sc.Session.Get(USERMODEL + getDate)
		if sessionStatus == nil {
			result = sc.Service.GetUserDailyCount(getDate)
			sc.Session.Set(USERMODEL+getDate, result)
		} else {
			sessionResult := sessionStatus.(int64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": util.RECODE_OK,
					"count":  sessionResult,
				},
			}
		}
	case "order":
		sessionStatus := sc.Session.Get(ORDERMODEL + getDate)
		if sessionStatus == nil {
			result = sc.Service.GetOrderDailyCount(getDate)
			sc.Session.Set(ORDERMODEL+getDate, result)
		} else {
			sessionResult := sessionStatus.(int64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": util.RECODE_OK,
					"count":  sessionResult,
				},
			}
		}

	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": util.RECODE_OK,
			"count":  result,
		},
	}
}
