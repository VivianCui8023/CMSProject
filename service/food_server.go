package service

import (
	"github.com/go-xorm/xorm"
	"webProject/model"
)

type FoodServer interface {
	//获取食品记录总数
	GetFoodCount() (int64, error)
}

type foodServer struct {
	engin *xorm.Engine
}

func NewFoodServer(db *xorm.Engine) FoodServer {
	return &foodServer{
		engin: db,
	}
}

func (fs *foodServer) GetFoodCount() (int64, error) {
	engin := fs.engin
	count, err := engin.Count(new(model.Food))
	return count, err
}
