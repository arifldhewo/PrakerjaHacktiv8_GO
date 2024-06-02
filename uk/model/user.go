package model

import "time"

type User struct {
	ID        uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Username  string `json:"username" gorm:"column:username;not null;unique"`
	Email     string `json:"email" gorm:"column:email;not null; unique"`
	Password  string `json:"password" gorm:"column:password;not null"`
	Age       uint8  `json:"age" gorm:"column:age;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt;not null"`
}