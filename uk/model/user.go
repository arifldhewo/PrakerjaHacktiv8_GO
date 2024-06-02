package model

import (
	"time"
)

type User struct {
	ID        uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Username  string `json:"username" gorm:"column:username;not null;unique"`
	Email     string `json:"email" gorm:"column:email;not null; unique"`
	Password  string `json:"password,omitempty" gorm:"column:password;not null;"`
	Age       uint8  `json:"age,omitempty" gorm:"column:age;not null"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:createdAt;not null"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt;not null"`
}