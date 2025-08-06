package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Initialize() {
	var err error
	DB, err = sql.Open("sqlite3", "app.db")

	if err != nil {
		fmt.Println("⚠️ sql open() error ==>", err)
		panic("Could not access to data base (create or open).")
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("⚠️ Ping() error ==>", err)
		panic("Could not connect to database.")
	}

	// defer DB.Close()

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createEventsTable()

	if err != nil {
		fmt.Println("⚠️ createEventsTable() error ==>", err)
		panic("Could not create events table.")
	}

	err = createUsersTable()
	if err != nil {
		fmt.Println("⚠️ createUsersTable() error ==>", err)
		panic("Could not create users table.")
	}
}

func createEventsTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			due_date DATETIME NOT NULL,
			create_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := DB.Exec(query)

	return err
}

func createUsersTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			first_name TEXT DEFAULT '',
			last_name TEXT DEFAULT ''
		)
	`

	result, err := DB.Exec(query)

	fmt.Println("result from create users table", result)
	return err
}
