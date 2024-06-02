package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uk/helper"
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