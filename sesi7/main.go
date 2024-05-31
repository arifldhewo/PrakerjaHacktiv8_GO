package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sesi6/controller"
	"github.com/sesi6/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
	
func main () {
	route := gin.Default()
	dbconf := "host=localhost user=postgres password=root dbname=ecommerce port=5432 TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dbconf), &gorm.Config{}) 

	if err != nil {
		log.Fatal(err)
		return
	}
	// Repo
	productRepo := &repository.ProductRepo{DB: db}

	// CI/CD
	productRepo.Migrate()

	// Controller
	productController := &controller.ProductController{Repository: productRepo}

	productGroup := route.Group("/v1")
	{
		productGroup.GET("/products", productController.GetGorm)
		productGroup.POST("/products", productController.CreateGorm)
		productGroup.PUT("/products/:id", productController.UpdateGorm)
		productGroup.DELETE("products/:id", productController.DeleteGorm)
	}

	err = route.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}