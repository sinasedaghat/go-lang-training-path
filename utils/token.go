package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "this-is-secret-key"

func GenerateToken(id int, email string) (string, error) {
	// 	claims := jwt.MapClaims{
	// 		"id": id,
	// 		"email": email,
	// 		"exp": time.Now().Add(30 * time.Minute).Unix(),
	// 	}

	// token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// return  token.SignedString(secretKey)

	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(30 * time.Minute).Unix(),
	}).SignedString([]byte(secretKey))
}

func ParseTokenAndGetID(token string) (bool, int) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signature method is incorrect")
		}

		return []byte(secretKey), nil
	})

	// TODO: return correct id
	if err != nil {
		fmt.Println("this error from ParseTokenAndGetID function in utils.token package", err)
		return false, 0
	}

	fmt.Println("parsedToken from ParseTokenAndGetID function in utils.token package", parsedToken)
	return true, 1000
}
