package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uk/controller"
	"github.com/uk/middleware"
	"github.com/uk/repository"
	"github.com/uk/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var engine *gin.Engine

func init() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("error: %v", err)
	}

	// Init DB
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	pwd	 := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable Timezone=Asia/Jakarta", host, port, user, name, pwd)
	dbLocal, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	db = dbLocal

	// Init Gin HTTP
	engine = gin.Default()
}

func main () {

	// Repository
	userRepo := &repository.UserRepo{DB: db}

	// Migrate
	userRepo.Migrate()

	// Service
	userService := &service.UserService{Repository: userRepo}

	// Controller
	userController := &controller.UserController{Service: userService}

	usersGroup := engine.Group("/users")
	{
		usersGroup.POST("/login", userController.Login)
		usersGroup.POST("/register", userController.Register)

		usersGroup.Use(middleware.BearerAuthorization())

		usersGroup.PUT("/:id", userController.Update)
		usersGroup.DELETE("/:id", userController.Delete)
	}

	if err := engine.Run(":8000"); err != nil {
		log.Fatal(err)
	}


}