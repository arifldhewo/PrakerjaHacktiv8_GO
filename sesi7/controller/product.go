package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sesi7/helper"
	"github.com/sesi7/model"
	"github.com/sesi7/repository"
)

type ProductController struct {
	Repository *repository.ProductRepo
}

func (p *ProductController) GetGorm(ctx *gin.Context) {
	products, err := p.Repository.Get()
	
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateGorm(ctx *gin.Context) {
	product := &model.Product{}

	headerAuth := ctx.GetHeader("Authorization")

	splitToken := helper.SplitJWT(headerAuth)

	if splitToken == "Error Authorization" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
			"message": "something went wrong",
		})
		return
	}

	valid, err := helper.ValidateUserJWT1(splitToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
			"message": "Invalid Token",
		})
		return
	}

	claims, err := helper.GetJWTPayload(valid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
			"message": "Invalid Token",
		})
		return
	}

	if err := ctx.Bind(product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "something went wrong",
		})
		log.Panicln(err)
		return
	}

	uId := uint64(claims["id"].(float64))

	product.UserID = uId

	errProd := p.Repository.Create(product)
	if errProd != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (p *ProductController) UpdateGorm(ctx *gin.Context) {
	product := &model.Product{}
	path := ctx.Param("id")

	if err := ctx.Bind(product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message" : "Bad Request",
		})
		return
	}

	err := p.Repository.Update(&path, product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message" : "something went wrong",
		})
		return
	}

	strpath, err := strconv.Atoi(path)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message" : "Convert Path Int to path is error",
		})
	}

	product.ID = uint64(strpath)

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "data updated",
		"data": product,
	})
}

func (p *ProductController) DeleteGorm(ctx *gin.Context) {
	product := &model.Product{}
	path := ctx.Param("id")

	err := p.Repository.Delete(&path, product)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message" : "something went wrong",
		})
		return
	}

	
	ctx.JSON(http.StatusNoContent, gin.H{
		
	})
}