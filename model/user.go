package model

import "time"

type User struct {
	Id       int64     `xorm:"pk autoincr" json:"id"`
	UserName string    `xorm:"varchar(12)" json:"username"`
	Register time.Time `json:"register_time"`
	Mobile   string    `xorm:"varchar(11)" json:"mobile"` //手机号是固定11位数字所以一般用char11存储
	IsActive int64     `json:"is_active"`
	Balance  int64     `json:"balance"`                    //用户的账户余额
	Avatar   string    `xorm:"varchar(255)" json:"avatar"` //头像信息
	Pwd      string    `json:"pwd"`
	DelFlag  int64     `json:"delFlag"`
	CityName string    `xorm:"varchar(24)" json:"city_name"`
}

func (user *User) UsertoResponeDec() interface{} {
	return map[string]interface{}{
		"id":            user.Id,
		"userName":      user.UserName,
		"register_time": user.Register,
		"mobile":        user.Mobile,
		"is_active":     user.IsActive,
		"balance":       user.Balance,
		"avatar":        user.Avatar,
		"delFlag":       user.DelFlag,
		"city":          user.CityName,
	}
}
