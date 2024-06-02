package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"
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
	if errMessage, err := service.ValidateTitleCaptionPhoto(photo); err != nil {
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

func (p *PhotoController) Update(ctx *gin.Context) {
	photo := &model.Photo{}

	ctx.Bind(&photo)

	// Validation For Photo Input
	if errMessage, err := service.ValidateTitleCaptionPhoto(photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errMessage)
		return
	}

	// Get id photo from param
	idStr := ctx.Param("id")

	// Convert String Id photo
	uId, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	// Update Data
	if err := p.Service.Update(uId, photo); err != nil {
		if strings.Contains(err.Error(), "photo_url") {
			err = errors.New("photo_url already taken")
		}
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error" : err.Error(),
		})
		return
	}

	updatedPhoto := model.Photo{}

	// Get Updated Data
	if err := p.Service.GetPhotoById(uId, &updatedPhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : "something went wrong",
		})
		return
	}

	// Return the updatedPhoto
	ctx.JSON(http.StatusCreated, gin.H{
		"id" : updatedPhoto.ID,
		"title": updatedPhoto.Title,
		"caption": updatedPhoto.Caption,
		"photo_url": updatedPhoto.PhotoUrl,
		"user_id": updatedPhoto.UserID,
		"updatedAt": updatedPhoto.UpdatedAt,
	})
}

func (p *PhotoController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : "something went wrong1",
		})
		return
	}

	if err := p.Service.Delete(idInt); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been deleted",
	})
}