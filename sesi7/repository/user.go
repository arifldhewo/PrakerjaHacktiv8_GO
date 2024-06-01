package repository

import (
	"log"

	"github.com/sesi7/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) Migrate() {
	err := u.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}
}

func (u *UserRepo) Create(user *model.User) error {
	err := u.DB.Debug().Model(&model.User{}).Create(user).Error
	return err
}

func (u *UserRepo) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := u.DB.Debug().Model(&model.User{}).Where("email = ?", email).First(&user).Error
	return user, err
}