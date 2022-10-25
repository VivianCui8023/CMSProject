package model

type food struct {
	Id          int           `json:"item_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	MonthSales  int           `json:"month_sales"`
	Rating      int           `json:"rating"`
	ImagePath   string        `json:"image_path"`
	Activity    string        `json:"activity"`
	Attributes  string        `json:"attributes"`
	Specs       string        `json:"specs"`
	CategoryId  int64         `xorm:"index"`
	Category    *FoodCategory `xorm:"-"`
	Restaurant  *Shop         `xorm:"-"`
	DelFlag     int           `json:"del_flag"`
}
