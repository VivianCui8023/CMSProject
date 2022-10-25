package model

type FoodCategory struct {
	Id               int64  `xorm:"pk autoincr" json:"id"`
	CategoryName     string `json:"name"`
	CategoryDec      string `json:"description"`
	Level            int64  `json:"level"`
	ParentCategoryId int64  `json:"parent_category_id"`
	RestaurantId     int64  `xorm:"index" json:"restaurant_id"`
}
