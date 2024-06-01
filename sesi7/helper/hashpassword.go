package helper

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	SECRET_KEY []byte = []byte("20242213960490267598567908189775563674735941821979632997103732363925894510946901202218157")
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUserJWT(email string, id uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id": id,
			"email":    email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateUserJWT(token string) bool {
	jwttoken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})
	if err != nil {
		return false
	}
	return jwttoken.Valid
}

func SplitJWT(token string) string {
	splitToken := strings.Split(token, " ")
	if len(splitToken) != 2{
		return "Error Authorization"
	} else if splitToken[0] != "Bearer"{
		return "Error Authorization"
	}
	return splitToken[1]
}

func ValidateUserJWT1(token string) (*jwt.Token, error) {
	jwttoken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return jwttoken, err
}

func GetJWTPayload(token *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid Token")
}