package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
	"sample-url.com/REST-API/models"
)

func createUsers(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println("error from create user read body of request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "can't pars request's body"})
		return
	}

	err = user.Save()

	if err != nil {
		sqliteErr, ok := err.(sqlite3.Error)

		fmt.Println("sqliteErr ====>", sqliteErr)
		fmt.Println("ok ===> ", ok)
		fmt.Println("sqliteErr.ExtendedCode ===> ", sqliteErr.ExtendedCode)
		fmt.Println("sqlite3.ErrConstraintUnique ===> ", sqlite3.ErrConstraintUnique)

		fmt.Println("error from create user save function", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something is wrong"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "create new user successful"})
}
