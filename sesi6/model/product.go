package model

var Products = []*Product{}

type Product struct {
	ID    uint64  `json:"id" gorm:"primary_key;auto_increment"`
	Name  string  `json:"name" gorm:"column:name"`
	Price float64 `json:"price" gorm:"column:price"`
}
