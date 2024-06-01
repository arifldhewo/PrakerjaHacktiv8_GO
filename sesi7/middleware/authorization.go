package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sesi7/helper"
	"github.com/sesi7/model"
	"gorm.io/gorm"
)

func BearerAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")
		// {Authorization: Bearer jwt_token}
		// get the encoded string
		splitToken := strings.Split(headerAuth, " ")
		if len(splitToken) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization header",
			})
			return
		}

		// check basic
		if splitToken[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid authorization method",
			})
			return
		}
		// validate jwt
		valid := helper.ValidateUserJWT(splitToken[1])
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "malformed token",
			})
			return
		}
		ctx.Next()
	}
}

func IsUser(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context)  {
		headerAuth := ctx.GetHeader("Authorization")

		splitToken := helper.SplitJWT(headerAuth)
		if splitToken == "Error Authorization" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "malformed token",
			})
			return
		}

		valid, err := helper.ValidateUserJWT1(splitToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "malformed token",
			})
			return
		}

		claims, err := helper.GetJWTPayload(valid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "invalid token",
			})
			return
		}

		path := ctx.Param("id")

		product := &model.Product{}

		errFetchProd := DB.Model(&model.Product{}).Where("id = ?", path).First(&product).Error
		if errFetchProd != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
				"message": "something went wrong",
			})
			return
		}

		for key, val := range claims{
			if key == "id"{
				if uint64(val.(float64)) == product.UserID{
					ctx.Next()
				} else {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
						"message": "You're not the owner of the product",
					})
					return
				}
			}
		}

	}
}