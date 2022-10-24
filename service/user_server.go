package service

import (
	"github.com/go-xorm/xorm"
	"webProject/model"
)

type UserService interface {
	GetUserCount() (int64, error)
	GetUserList(offset int, limit int) []model.User
}

func NewUserService(db *xorm.Engine) UserService {
	return &userService{
		engin: db,
	}
}

// 实现类
type userService struct {
	engin *xorm.Engine
}

//实现接口函数

func (us *userService) GetUserCount() (int64, error) {
	engine := us.engin
	result, err := engine.Where("del_flag = ?", 0).Count(model.User{})
	if err != nil {
		panic(err.Error())
		return 0, err
	}
	return result, nil
}

// 查询用户列表
func (us *userService) GetUserList(offset int, limit int) []model.User {
	result := make([]model.User, 0)
	engin := us.engin

	err := engin.Limit(limit, offset*limit).Find(&result)
	if err != nil {
		panic(err.Error())
	}
	return result
}
