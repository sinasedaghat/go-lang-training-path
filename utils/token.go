package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("this-is-secret-key")

// func GenerateToken(id int, email string) (string, error) {
// 		claims := jwt.MapClaims{
// 			"id": id,
// 			"email": email,
// 			"exp": time.Now().Add(30 * time.Minute).Unix(),
// 		}
// 	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
// 	return  token.SignedString(secretKey)
// }

func GenerateToken(id int, email string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(30 * time.Minute).Unix(),
	}).SignedString(secretKey)
}

func tokenParser(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signature method is incorrect")
		}

		return secretKey, nil
	})
}

func authenticationFailed() (int, error) {
	return 0, errors.New("authentication is failed")
}

func ParseTokenAndGetID(token string) (int, error) {
	parsedToken, err := tokenParser(token)

	if err != nil {
		fmt.Println("🔑 error from jwt.Parse(...) function: ", err)
		return authenticationFailed()
	}

	if validate := parsedToken.Valid; !validate {
		fmt.Println("🔑 this error when validate parsed token from ParseTokenAndGetID function in utils.token package")
		return authenticationFailed()
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("🔑 this error from check type claims in ParseTokenAndGetID function in utils.token package")
		return authenticationFailed()
	}

	id, _ := claims["id"].(float64)
	userId := int(id)

	return userId, nil
}

func ValidateAndParseToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := tokenParser(token)

	if err != nil {
		fmt.Println("🔑 error from jwt.Parse(...) function: ", err)
		return nil, errors.New("authentication failed")
	}

	if climes, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return climes, nil
	}

	fmt.Println("🔑 climes is not correct, or parsedToken is not valid")
	return nil, errors.New("authentication failed")
}
