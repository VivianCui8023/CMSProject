package model

import "time"

/*
model
*/
// 管理员用户信息

type Admin struct {
	AdminId   int64     `xorm:"pk autoincr" json:"id"`
	AdminName string    `xorm:"varchar(32)" json:"admin_name"`
	CreatTime time.Time `xorm:"DateTime" json:"creat_time"`
	Status    int64     `xorm:"default 0" json:"status"`
	Avatar    string    `xorm:"varchar(255)" json:"avatar"`
	Pwd       string    `xorm:"varchar(255)" json:"pwd"`
	CityName  string    `xorm:"varchar(22)" json:"city_name"`
	CityId    int64     `xorm:"index" json:"city_id"`
	//City      *City     `xorm:"- <- ->"`
}

// type City struct {
// }
// 将admin转化为json形式

func (ad *Admin) AdmintoRespone() interface{} {
	respDesc := map[string]interface{}{
		"user_name": ad.AdminName,
		"id":        ad.AdminId,
		"creatTime": ad.CreatTime,
		"status":    ad.Status,
		"avatar":    ad.Avatar,
		"city":      ad.CityName,
		"admin":     "管理员",
	}
	return respDesc
}
