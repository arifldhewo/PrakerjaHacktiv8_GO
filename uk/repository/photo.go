package repository

import (
	"log"

	"github.com/uk/model"
	"gorm.io/gorm"
)

type PhotoRepo struct {
	DB *gorm.DB
}

func (p *PhotoRepo) Migrate() {
	if err := p.DB.AutoMigrate(&model.Photo{}); err != nil {
		log.Fatalf("Failed to perform migration: %v", err)
	}
}

func (p *PhotoRepo) Create(photo *model.Photo) error {
	if err := p.DB.Model(&model.Photo{}).Create(&photo).Error; err != nil {
		return err
	}
	return nil
}