package model

type User struct {
	ID       uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}