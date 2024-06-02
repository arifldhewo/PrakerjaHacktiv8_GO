package service

import (
	"github.com/asaskevich/govalidator"
	"github.com/uk/helper"
	"github.com/uk/model"
	"github.com/uk/repository"
	"gorm.io/gorm"
)

type PhotoService struct {
	Repository *repository.PhotoRepo
}

func ValidateCreatePhoto(input *model.Photo) (map[string]any, error) {
	errMessage := make(map[string]any)
	requirement := map[string]any{
		"title":    "required",
		"caption": "required",
		"photo_url": "required",
	}

	inputMap := map[string]any {
		"title" : input.Title,
		"caption": input.Caption,
		"photo_url": input.PhotoUrl,
	}

	result, err := govalidator.ValidateMap(inputMap, requirement)
	if err != nil || !result {
		errMessage = helper.ValidatorErrExtractor(err.Error())
		return errMessage, err 
	}
	return errMessage, err 
}

func (p *PhotoService) GetUserLastPhoto(userId int, photo *model.Photo) error {
	if err := p.Repository.DB.Model(&model.Photo{}).Where("user_id = ?", userId).Last(&photo).Error; err != nil {
		return err
	}
	return nil
}

func (p *PhotoService) GetAllPhotos(photos *[]model.Photo) error {

	if err := p.Repository.DB.Preload("User", func (db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "email")
	}).Find(&photos).Error; err != nil {
		return err
	}
	return nil
}