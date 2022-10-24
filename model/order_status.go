package model

/*
*
  - 订单状态结构体定义（数据库中定义好）

1	未支付
2	已支付
3	已发货
4	正在配送
5	已接收
6	发起退款
7	正在退款
8	取消订单
*/
type OrderStatus struct {
	//Id         int    `xorm:"pk autoincr" json:"id"` //主键
	StatusId   int    //订单状态编号
	StatusDesc string `xorm:"varchar(255)"` // 订单状态描述

}
