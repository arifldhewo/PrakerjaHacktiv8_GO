package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uk/helper"
	"github.com/uk/model"
	"github.com/uk/service"
)

type UserController struct {
	Service *service.UserService
}

func (u *UserController,) Register(ctx *gin.Context) {

	user := &model.User{}

	ctx.Bind(&user)

	// * Validate The Request
	if errMessage, err := service.ValidateRegister(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errMessage)
		return
	}

	// * Hashing Password
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message" : "Something went wrong",
		})
		return
	}

	// * Reassign password -> hashed password
	user.Password = string(hashedPassword)

	// * Creating new data
	if err := u.Service.OneOfTheFieldAlreadyTaken(user, "POST"); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	// * New Copy Model For Responses
	submittedData := model.User{}

	// * Get the data and bind to submittedData model
	if err := u.Service.Repository.GetUserByEmail(user.Email, &submittedData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message" : "Something went wrong",
		})
		return
	}

	// * Give Response
	ctx.JSON(http.StatusCreated, gin.H{
		"age": submittedData.Age,
		"email": submittedData.Email,
		"id": submittedData.ID,
		"username": submittedData.Username,
	})
}

func (u *UserController) Login(ctx *gin.Context) {
	user := &model.User{}

	ctx.Bind(&user)

	// * Validate The Request
	if errMessage, err := service.ValidateEmailPassword(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errMessage)
		return
	}

	// * Fetch User Login Detail
	userFetch := model.User{}
	if err := u.Service.Repository.GetUserByEmail(user.Email, &userFetch); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Your email or password is wrong",
		})
		return
	}

	// * Comparing User Password
	if err := helper.ComparePassword(userFetch.Password, user.Password); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Your email or password is wrong",
		})
		return
	}

	// * Generate JWT
	tokenString, err := helper.GenerateJWT(&userFetch)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token" : tokenString,
	})
}

func (u *UserController) Update(ctx *gin.Context) {
	user := &model.User{}

	// * Get ID Param
	idStr := ctx.Param("id")

	ctx.Bind(&user)

	// * Convert to int
	uId, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	// * Validation for input
	if errMessage, err := service.ValidateEmailUsername(user); err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errMessage)
		return
	}

	// * Update User Data
	if err := u.Service.OneOfTheFieldAlreadyTaken(user, "PUT", uId); err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// * Iniate New Mode for response
	updateUser := model.User{}

	// * Bind data into updateUser
	if err := u.Service.Repository.GetUserByEmail(user.Email, &updateUser); err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	// * Return response updataUser
	ctx.JSON(http.StatusOK, gin.H{
		"id": updateUser.ID,
		"email": updateUser.Email,
		"username": updateUser.Username,
		"age": updateUser.Age,
		"updatedAt": updateUser.UpdatedAt,
	})
}

func (u *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	uId, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	if err := u.Service.Repository.DB.Model(&model.User{}).Where("id = ?", uId).Delete(&uId).Error; err != nil {
		log.Print(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}