package repository

import (
	"log"

	"github.com/sesi7/model"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func (p ProductRepo) Migrate() {
	err := p.DB.AutoMigrate(&model.Product{})
	if err != nil {
		log.Fatal(err)
	}
}

func (p *ProductRepo) Get() ([]*model.Product, error) {
	products := []*model.Product{}
	err := p.DB.Debug().Model(&model.Product{}).Find(&products).Error
	return products, err
}

func (p *ProductRepo) Create(product *model.Product) error {
	err := p.DB.Debug().Model(&model.Product{}).Create(product).Error
	return err
}

func (p *ProductRepo) Update(path *string, product *model.Product) error {
	err := p.DB.Debug().Model(&model.Product{}).Where("id = ?", path).Updates(product).Error
	return err
}

func (p *ProductRepo) Delete(path *string, product *model.Product) error {
	err := p.DB.Debug().Model(&model.Product{}).Delete(product, path).Error
	return err
}
