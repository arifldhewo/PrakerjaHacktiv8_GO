package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sesi7/helper"
	"github.com/sesi7/model"
	"github.com/sesi7/repository"
)

type UserController struct {
	Repository *repository.UserRepo
}

func (u *UserController) Registration(ctx *gin.Context) {
	user := &model.User{}

	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}

	user.Password = hashPassword

	err = u.Repository.Create(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User Registration successfully",
	})
}

func (u *UserController) Login(ctx *gin.Context) {
	user := &model.User{}

	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}

	userFetched, err := u.Repository.GetByEmail(user.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "user not found",
		})
		return
	}

	valid := helper.CheckPasswordHash(user.Password, userFetched.Password)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
			"message": "wrong password",
		})
		return
	}

	token, err := helper.GenerateUserJWT(userFetched.Email, userFetched.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}