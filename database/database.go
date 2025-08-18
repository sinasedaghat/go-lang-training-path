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
		fmt.Println("🗄️ sql.open() has  error:", err)
		panic("Could not access to data base (create or open).")
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("🗄️ Ping() has error:", err)
		panic("Could not connect to database.")
	}

	// defer DB.Close()

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createUsersTable()
	if err != nil {
		fmt.Println("🗄️ createUsersTable() has error:", err)
		panic("Could not create users table.")
	}

	err = createEventsTable()
	if err != nil {
		fmt.Println("🗄️ createEventsTable() has error:", err)
		panic("Could not create events table.")
	}

	err = createRegistrationTable()
	if err != nil {
		fmt.Println("🗄️ createRegistrationTable() has error:", err)
		panic("Could not create registration table.")
	}
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
	_, err := DB.Exec(query)

	return err
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
			create_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
	_, err := DB.Exec(query)

	return err
}

func createRegistrationTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS registration(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`
	_, err := DB.Exec(query)

	return err
}
