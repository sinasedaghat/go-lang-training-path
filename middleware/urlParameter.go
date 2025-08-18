package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HINT: I say, this is exactly a `gin.HandlerFunc`; slightly more verbose.
func URLParameter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			fmt.Println("🔀 Error extracting ID from URL:", err)
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "The desired item was not found."})
			return
		}

		ctx.Set("id_param", id)
		ctx.Next()
	}
}
