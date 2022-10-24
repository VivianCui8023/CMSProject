package model

/**
 * 地址结构实体
 */
type Address struct {
	ID            int64  `xorm:"pk autoincr" json:"id"`
	AddressId     int    `xorm:"index" json:"address_id"`
	Address       string `xorm:"varchar(255)" json:"address"`
	AddressDetail string `xorm:"varchar(255)" json:"address_detail"`
	IsValid       int    `json:"is_valid"`
}
