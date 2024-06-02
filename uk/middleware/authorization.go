package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uk/helper"
	"github.com/uk/model"
	"gorm.io/gorm"
)

func BearerAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")

		claims, err := helper.GetJWTPayload(headerAuth)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		} 

		log.Print(claims)
		ctx.Next()
	}
}

func IsUser(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")

		claims, err := helper.GetJWTPayload(headerAuth)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not the owner of this resource",
			})
			return
		}

		uId := int(claims["user_id"].(float64))
		idStr := ctx.Param("id")

		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong",
			})
			return
		}

		photo := &model.Photo{}

		if err := DB.Model(&model.Photo{}).Where("id = ?", idInt).First(&photo).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data not found",
			})
			return
		}

		if photo.UserID != uint64(uId) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not the owner of this resource",
			})
			return
		} else {
			ctx.Next()
		}
	}
}