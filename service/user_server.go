package service

import (
	"github.com/go-xorm/xorm"
	"webProject/model"
)

type UserService interface {
	GetUserCount() int64
}

// 实现类
type userService struct {
	engin xorm.Engine
}

//实现接口函数

func (us *userService) GetUserCount() int64 {
	engine := us.engin
	result, err := engine.Count(model.User{})
	if err != nil {
		panic(err.Error())
		return 0
	}
	return result
}
