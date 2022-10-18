package service

import (
	"github.com/go-xorm/xorm"
	"time"
	"webProject/model"
)

type StatisServer interface {

	//查询某一天用户增长量,因为Count返回的是int64

	GetUserDailyCount(data string) int64

	//查询某一天订单增长量
	//GetOrderDailyCount(data string) int64

	//查询管理员增长数量
	GetAdminDailyCount(data string) int64
}

type statisService struct {
	engin *xorm.Engine
}

func NewStatisService(db *xorm.Engine) StatisServer {
	return &statisService{
		engin: db,
	}
}

func (ss *statisService) GetAdminDailyCount(data string) int64 {
	engin := ss.engin

	startData, err := time.Parse("2006-01-02", data)
	if err != nil {
		return 0
	}
	endData := startData.AddDate(0, 0, 1)

	result, err := engin.Where("creat_time between ? and ? and status = 0",
		startData.Format("2006-01-02 15:04:05"),
		endData.Format("2006-01-02 15:04:06")).Count(model.Admin{})
	if err != nil {
		return 0
	}
	return result
}

func (ss *statisService) GetUserDailyCount(date string) int64 {
	engin := ss.engin

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	endDate := startDate.AddDate(0, 0, 1)

	result, err := engin.Where("register between ? and ? and del_flag = 0",
		startDate.Format("2006-01-02 15:04:05"),
		endDate.Format("2006-01-02 15:04:06")).Count(model.User{})

	if err != nil {
		return 0
	}
	return result
}

//func (ss *statisService) GetOrderDailyCount(data string) int64 {
//
//}
