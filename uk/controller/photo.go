package controller

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uk/helper"
	"github.com/uk/model"
	"github.com/uk/service"
)

type PhotoController struct {
	Service *service.PhotoService
}

func (p *PhotoController) Create(ctx *gin.Context) {
	photo := &model.Photo{}
	headerAuth := ctx.GetHeader("Authorization")

	ctx.Bind(photo)

	// Validation For Photo Input
	if errMessage, err := service.ValidateCreatePhoto(photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errMessage)
		return
	}

	// Get JWT Payload
	claims, err := helper.GetJWTPayload(headerAuth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	// Insert user_id from JWT token to model photo
	photo.UserID = uint64(claims["user_id"].(float64))

	// Create User Photo
	if err := p.Service.Repository.Create(photo); err != nil {
		if strings.Contains(err.Error(), "photo_url") {
			err = errors.New("photo_url already taken")
		}
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error" : err,
		})
		return
	}

	// Make empty model
	lastPhotoFromUser := model.Photo{}

	// Bind the data into lastPhotoFromUser struct
	if err := p.Service.GetUserLastPhoto(int(photo.UserID), &lastPhotoFromUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	// Return the lastPhotoFromUser
	ctx.JSON(http.StatusCreated, gin.H{
		"id" : lastPhotoFromUser.ID,
		"title": lastPhotoFromUser.Title,
		"caption": lastPhotoFromUser.Caption,
		"photo_url": lastPhotoFromUser.PhotoUrl,
		"user_id": lastPhotoFromUser.UserID,
		"createdAt": lastPhotoFromUser.CreatedAt,
	})
}

func (p *PhotoController) Get(ctx *gin.Context) {
	photos := []model.Photo{}

	if err := p.Service.GetAllPhotos(&photos); err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}