package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sample-url.com/REST-API/utils"
)

func Authentication(ctx *gin.Context) {
	// token := ctx.Request.Header.Get("Authentication")
	token := ctx.GetHeader("Authentication")

	if token == "" {
		fmt.Println("🔀 header don't have token")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please signIn."})
	}

	claim, err := utils.ValidateAndParseToken(token)

	if err != nil {
		fmt.Println("🔀 error after validate and parse token: ", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed!"})
	}

	id, _ := claim["id"].(float64)
	userId := int(id)

	ctx.Set("user_id", userId)
	ctx.Next()
}
