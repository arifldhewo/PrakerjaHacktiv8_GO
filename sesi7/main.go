package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sesi7/controller"
	"github.com/sesi7/middleware"
	"github.com/sesi7/repository"
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
	userRepo := &repository.UserRepo{DB: db}
	productRepo := &repository.ProductRepo{DB: db}

	// CI/CD
	userRepo.Migrate()
	productRepo.Migrate()


	// Controller
	userController := &controller.UserController{Repository: userRepo}
	productController := &controller.ProductController{Repository: productRepo}

	Group := route.Group("/v1")
	{
		Group.POST("/registrations", userController.Registration)
		Group.POST("/logins", userController.Login)

		Group.Use(
			middleware.BearerAuthorization(),
		)
		Group.GET("/products", productController.GetGorm)
		Group.POST("/products", productController.CreateGorm)
		Group.Use(
			middleware.IsUser(db),
		)
		Group.PUT("/products/:id", productController.UpdateGorm)
		Group.DELETE("products/:id", productController.DeleteGorm)
	}

	err = route.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}