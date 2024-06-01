package model

var Products = []*Product{}

type Product struct {
	ID     uint64  `json:"id" gorm:"primary_key;auto_increment"`
	UserID uint64  `json:"user_id" gorm:"column:user_id"`
	Name   string  `json:"name" gorm:"column:name"`
	Price  float64 `json:"price" gorm:"column:price"`
}
