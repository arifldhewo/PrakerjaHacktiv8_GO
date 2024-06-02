package repository

import (
	"log"

	"github.com/uk/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) Migrate() {
	err := u.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to perform migrations: %v", err)
	}
}

func (u *UserRepo) GetUserByEmail(email string, user *model.User) error {
	err := u.DB.Model(&model.User{}).Where("email = ?", email).First(&user).Error
	return err
}

func (u *UserRepo) Update(id int, user *model.User) error {
	err := u.DB.Model(&model.User{}).Where("id = ?", id).Updates(&user).Error
	return err
}