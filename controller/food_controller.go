package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"webProject/service"
	"webProject/util"
)

type FoodController struct {
	Server service.FoodServer
	Ctx    iris.Context
}

//获取食品记录总数 get  /foods/count

func (fc *FoodController) GetCount() mvc.Response {
	count, err := fc.Server.GetFoodCount()
	if err != nil {
		iris.New().Logger().Error(err.Error())
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
			"count":  count,
		},
	}
}