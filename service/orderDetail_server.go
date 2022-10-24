package service

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"

	"webProject/model"
)

type OrderDetailServer interface {
	GetOrderList(offset, limit int) []model.OrderDetail
}

type orderDetailServer struct {
	engin *xorm.Engine
}

func NewOrderDetailServer(db *xorm.Engine) OrderDetailServer {
	return &orderDetailServer{
		engin: db,
	}
}

func (ods *orderDetailServer) GetOrderList(offset, limit int) []model.OrderDetail {
	orderList := make([]model.OrderDetail, 0)

	err := ods.engin.Table("user_order").
		Join("INNER", "order_status", "order_status.status_id = user_order.order_status_id").
		Join("INNER", "user", "user_order.user_id = user.id").
		Join("INNER", "shop", "user_order.shop_id = shop.id").
		Join("INNER", "address", "user_order.address_id = address.id").
		Limit(limit, limit*offset).Find(&orderList)

	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}

	return orderList
}
