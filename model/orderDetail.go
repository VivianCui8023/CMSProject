package model

/*
*订单查询信息
 */

type OrderDetail struct {
	//这个顺序要跟表联查顺序一致
	UserOrder   `xorm:"extends"`
	User        `xorm:"extends"`
	OrderStatus `xorm:"extends"`
	Shop        `xorm:"extends"`
	Address     `xorm:"extends"`
}

func (od *OrderDetail) OrderDetail2Resp() interface{} {
	resDesc := map[string]interface{}{
		"id":                   od.UserOrder.Id,
		"total_amount":         od.UserOrder.SumMoney,
		"user_id":              od.User.UserName,
		"status":               od.OrderStatus.StatusDesc,
		"restaurant_id":        od.Shop.Id,
		"restaurant_image_url": od.Shop.ImagePath,
		"restaurant_name":      od.Shop.Name,
		"formatted_creat_at":   od.Time,
		"status_code":          0,
		"address_id":           od.Address.AddressId,
	}
	statusDesc := map[string]interface{}{
		"color":     "f60",
		"sub_title": "15分钟内支付",
		"title":     "title",
	}
	resDesc["status_bar"] = statusDesc
	return resDesc
}
