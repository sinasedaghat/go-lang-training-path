package utils

import (
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
