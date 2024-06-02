package model

import (
	"time"
)

type Photo struct {
	ID        uint64 `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null"`
	Title     string `json:"title" gorm:"column:title;not null"`
	Caption   string `json:"caption" gorm:"column:caption;not null"`
	PhotoUrl string `json:"photo_url" gorm:"column:photo_url;not null;unique"`
	UserID   uint64 `json:"user_id" gorm:"column:user_id"`
	User User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;not null"`
	UpdatedAt time.Time `Json:"updatedAt" gorm:"column:updatedAt;not null"`
}