package service

import (
	"errors"
	"log"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/uk/helper"
	"github.com/uk/model"
	"github.com/uk/repository"
)

type UserService struct {
	Repository *repository.UserRepo
}

func ValidateRegister(input *model.User) (map[string]any, error) {
	errMessage := make(map[string]any)
	requirement := map[string]any{
		"email":    "required,email",
		"username": "required,alphanum",
		"password": "required",
		"age":      "required,numeric",
	}

	inputMap := map[string]any {
		"email" : input.Email,
		"username": input.Username,
		"password": input.Password,
		"age": input.Age,
	}

	result, err := govalidator.ValidateMap(inputMap, requirement)
	if err != nil || !result {
		errMessage = helper.ValidatorErrExtractor(err.Error())
		return errMessage, err 
	}
	return errMessage, err 
}

func ValidateEmailPassword(input *model.User) (map[string]any, error) {
	errMessage := make(map[string]any)
	requirement := map[string]any{
		"email":    "required,email",
		"password": "required",
	}

	inputMap := map[string]any {
		"email" : input.Email,
		"password": input.Password,
	}

	result, err := govalidator.ValidateMap(inputMap, requirement)
	if err != nil || !result {
		errMessage = helper.ValidatorErrExtractor(err.Error())
		log.Printf("error: %v", err)
		return errMessage, err 
	}
	return errMessage, err 
}

func ValidateEmailUsername(input *model.User) (map[string]any, error) {
	errMessage := make(map[string]any)
	requirement := map[string]any{
		"email":    "required,email",
		"username": "required",
	}

	inputMap := map[string]any {
		"email" : input.Email,
		"username": input.Username,
	}

	result, err := govalidator.ValidateMap(inputMap, requirement)
	if err != nil || !result {
		errMessage = helper.ValidatorErrExtractor(err.Error())
		log.Printf("error: %v", err)
		return errMessage, err 
	}
	return errMessage, err 
}

func (u *UserService) OneOfTheFieldAlreadyTaken(user *model.User, method string, uId ...int) (error) {
	
	var err error

	switch method {
	case "POST":
		err = u.Repository.DB.Model(&model.User{}).Create(&user).Error
	
	case "PUT": 
		err = u.Repository.Update(uId[0], user)
	}

	if err != nil {
		if strings.Contains(err.Error(), "email") {
			return errors.New("email already taken")
		} else  {
			log.Print("masuk username")
			return errors.New("username already taken")
		}
	}
	return nil
}