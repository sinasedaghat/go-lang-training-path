package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
	"sample-url.com/REST-API/models"
	"sample-url.com/REST-API/utils"
)

func createUsers(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println("🫴 error from read body of request in the createUsers handler:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't pars request's body"})
		return
	}

	err = user.Save()

	if err != nil {
		// TODO: use these for create clear error
		sqliteErr, ok := err.(sqlite3.Error)
		fmt.Println("📋 (for TODO) sqliteErr ====>", sqliteErr)
		fmt.Println("📋 (for TODO) ok ===> ", ok)
		fmt.Println("📋 (for TODO) sqliteErr.ExtendedCode ===> ", sqliteErr.ExtendedCode)
		fmt.Println("📋 (for TODO) sqlite3.ErrConstraintUnique ===> ", sqlite3.ErrConstraintUnique)

		fmt.Println("🫴 error from user.save() in the createUsers handler:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something is wrong"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "create new user successful"})
}

func signInUsers(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println("🫴 error from read body of request in the signInUsers handler:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't pars request's body"})
		return
	}

	err = user.Get()

	if err != nil {
		fmt.Println("error from user.Get in user router file", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Email or Password is Invalid."})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		fmt.Println("🫴 error from utils.GenerateToken(...data) in the signInUsers handler:", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Email or Password is Invalid."})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "You are Authorize", "token": token})
}
