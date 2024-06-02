package helper

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/uk/model"
	"golang.org/x/crypto/bcrypt"
)

func ValidatorErrExtractor(err string) map[string]any {

	isMoreThanOneErr := strings.ContainsAny(err, ";")
	var splitErr []string
	mapVal := make(map[string]any)

	if isMoreThanOneErr {
		splitErr = strings.Split(err, ";")

		for key := range splitErr {
			tmp := strings.Split(splitErr[key], ":")
			
			for i := 0; i < len(tmp); i++ {
				if i + 1 == len(tmp) {
					break
				}

				if i %  2 == 0{
					mapVal[tmp[i]] = tmp[i+1]
				}
			}
		}

	} else {
		splitErr = strings.Split(err, ":")
		mapVal[splitErr[0]] = splitErr[1]
	}

	return mapVal
}

func ValidateRequest(model interface{}, ctx *gin.Context) (map[string]any, error ) {
	validate, err := govalidator.ValidateStruct(model)
	errMessage := make(map[string]any)
	if err != nil || !validate {
		errMessage := ValidatorErrExtractor(err.Error())
		return errMessage, err 
	}
	return errMessage, err 
}

func HashPassword(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error: %v", err)
	}
	return hashed, err
}

func ComparePassword(hashedPassword string, inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	if err != nil {
		log.Printf("error: %v", err)
	}
	return err
}

func GenerateJWT(u *model.User) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("error: %v", err)
	}

	SECRET_KEY := os.Getenv("JWT_SECRET_KEY")
	conv := []byte(SECRET_KEY)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"email": u.Email,
		"age": u.Age,
		"createdAt": u.CreatedAt,
		"updatedAt": u.UpdatedAt,
	})

	tokenString, err := token.SignedString([]byte(conv))
	if err != nil {
		log.Print(err)
		return "error: ", err
	}

	return tokenString, nil
}

func SplitJWT(token string) (string, error) {
	splitToken := strings.Split(token, " ")
	if len(splitToken) != 2{
		return "", errors.New("token malformed")
	} else if splitToken[0] != "Bearer"{
		return "", errors.New("error authorization method")
	}
	return splitToken[1], nil
}

func GetJWTPayload(token string) (jwt.MapClaims, error) {
	// Load env files
	if err := godotenv.Load(); err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	// Split Incoming token
	splitToken, err := SplitJWT(token)
	if err != nil {
		return nil, err
	}

	// Init Token
	SECRET_KEY := os.Getenv("JWT_SECRET_KEY")

	// Convert Token
	conv := []byte(SECRET_KEY)

	// Parse JWT Token
	jwtoken, err := jwt.Parse(splitToken, func(t *jwt.Token) (interface{}, error) {
		return conv, nil
	})

	if err != nil {
		return nil, errors.New("error: invalid token")
	}

	if claims, ok := jwtoken.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}