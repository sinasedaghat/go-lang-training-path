package models

import (
	"errors"
	"fmt"

	"sample-url.com/REST-API/database"
	"sample-url.com/REST-API/utils"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	RoleId    int    `json:"role_id" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// func (u User) claims() map[string]any {
// 	return map[string]any{
// 		"id":         u.ID,
// 		"email":      u.Email,
// 		"first_name": u.FirstName,
// 		"last_name":  u.LastName,
// 	}
// }

// 	claims := jwt.MapClaims{
// 		"id": id,
// 		"email": email,
// 		"exp": time.Now().Add(30 * time.Minute).Unix(),
// 	}

func (u User) Save() error {
	// query := `
	// 	INSERT INTO users(email, password, first_name, last_name)
	// 	VALUES (?, ?, COALESCE(?, ''), COALESCE(?, ''))
	// `

	query := `
		INSERT INTO users(email, password, role_id)
		VALUES (?, ?, ?)
	`

	statement, err := database.DB.Prepare(query)
	if err != nil {
		fmt.Println("🗂️ Error from Prepare(query) of user.Save method:", err)
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		fmt.Println("🗂️ Error from utils.HashPassword(password) of user.Save method:", err)
		return err
	}

	_, err = statement.Exec(u.Email, hashedPassword, u.RoleId)
	// result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		fmt.Println("🗂️ Error from Exec(data...) of user.Save method:", err)
	}

	return err
}

func (u *User) Get() error {
	query := `SELECT * FROM users WHERE email = ?`

	row := database.DB.QueryRow(query, u.Email)

	var password string

	row.Scan(&u.ID, &u.Email, &password, &u.FirstName, &u.LastName)

	if !utils.PasswordValidator(password, u.Password) {
		return errors.New("invalid password")
	}

	return nil
}
