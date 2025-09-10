package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/mattn/go-sqlite3"
	"sample-url.com/REST-API/constant"
	"sample-url.com/REST-API/models"
	"sample-url.com/REST-API/utils"
)

func createUser(ctx *gin.Context) {
	var user models.User
	user.RoleId = constant.DefaultRole()

	// override role from context if exists
	if value, exists := ctx.Get("role_id"); exists {
		if roleId, ok := value.(int); ok {
			user.RoleId = roleId
		}
	}

	// give data from body of request
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Printf("🫴 createUser() handler function has error when read body of request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request body"})
		return
	}

	// save user in DB
	if err := user.Save(); err != nil {
		// TODO: use these for create clear error
		// sqliteErr, ok := err.(sqlite3.Error)
		// fmt.Println("📋 (for TODO) sqliteErr ====>", sqliteErr)
		// fmt.Println("📋 (for TODO) ok ===> ", ok)
		// fmt.Println("📋 (for TODO) sqliteErr.ExtendedCode ===> ", sqliteErr.ExtendedCode)
		// fmt.Println("📋 (for TODO) sqlite3.ErrConstraintUnique ===> ", sqlite3.ErrConstraintUnique)

		// status, msg := utils.HandleSQLError(err)
		// ctx.JSON(status, gin.H{"success": false, "message": msg})

		log.Printf("🫴 createUser() handler function has error when user.Save() call: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "something is wrong"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "create new user successful"})
}

func signInUser(ctx *gin.Context) {
	var user models.User

	// give data from of request
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Printf("🫴 signInUser() handler function has error when read body of request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request body"})
		return
	}

	// get user from db
	if err := user.Get(); err != nil {
		log.Printf("🫴 signInUser() handler function has error when user.Get() call: %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Email or password is invalid"})
		return
	}

	// create JWT
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Printf("🫴 signInUser() handler function has error when utils.GenerateToken() call: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "You are Authorize",
		"data": gin.H{
			"token": token,
		},
	})
}

// func getUsers(ctx *gin.Context) {

// }
