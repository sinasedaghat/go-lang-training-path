package models

import (
	"fmt"

	"sample-url.com/REST-API/database"
	"sample-url.com/REST-API/utils"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u User) Save() error {
	// query := `
	// 	INSERT INTO users(email, password, first_name, last_name)
	// 	VALUES (?, ?, COALESCE(?, ''), COALESCE(?, ''))
	// `

	query := `
		INSERT INTO users(email, password)
		VALUES (?, ?)
	`

	statement, err := database.DB.Prepare(query)
	if err != nil {
		fmt.Println("⚠️ error from database.DB.Prepare() in save function ===>", err)
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		fmt.Println("⚠️ error from HashPassword function", err)
		return err
	}

	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		fmt.Println("⚠️ error from statement.Exec() in save function ===>", err)
	}

	fmt.Println("result from save user ===> ", result)
	return err
}
