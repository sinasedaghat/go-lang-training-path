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
		fmt.Println("🔀 Header don't have token")
		// ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Please signIn."})
		// ctx.Abort()
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please signIn."})
		return
	}

	claim, err := utils.ValidateAndParseToken(token)

	if err != nil {
		fmt.Println("🔀 Error after validate and parse token:", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed!"})
		return
	}

	id, _ := claim["id"].(float64)
	userId := int(id)

	ctx.Set("user_id", userId)
	ctx.Next()
}
