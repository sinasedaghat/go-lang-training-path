package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"sample-url.com/REST-API/constant"
)

var DB *sql.DB

func Initialize() {
	var err error
	// Open database connection
	DB, err = sql.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatalf("🗄️ sql.Open() error: %v", err)
	}

	// Ensure foreign key constraints are enforced
	if _, err := DB.Exec("PRAGMA foreign_keys = ON;"); err != nil { // Without this, the FK is ignored.
		log.Fatalf("🗄️ Enable foreign keys (DB.Exec(PRAGMA foreign_keys query) error: %v", err)
	}

	// Verify connection
	if err := DB.Ping(); err != nil {
		log.Fatalf("🗄️ Ping() error: %v", err)
	}

	//Connection pool tuning
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Migrations
	must(createUsersTable, "users")
	must(createEventsTable, "events")
	must(createRegistrationTable, "registration")
	must(createRolesTable, "roles")

	// Seeds
	/*
		when we have many seed functions;
		we can use this struct:
			[]struct {
				name string
				fn   func() error
			}
		and create struct of seeds and use for loop for call all of seed functions.
	*/
	if err := seedRoles(); err != nil {
		log.Fatalf("🌱 Failed to seed %s: %v", "roles", err)
	}
}

func must(createFunc func() error, tableName string) {
	if err := createFunc(); err != nil {
		log.Fatalf("🗄️ Failed to create %s table: %v", tableName, err)
	}
}

func createRolesTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS roles (
			id INTEGER PRIMARY KEY,
			label TEXT NOT NULL UNIQUE
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
			last_name TEXT DEFAULT '',
			role_id INTEGER,

			FOREIGN KEY (role_id)
				REFERENCES roles(id)
				ON DELETE RESTRICT
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

			FOREIGN KEY (user_id) 
				REFERENCES users(id)
				ON DELETE CASCADE
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

			FOREIGN KEY (event_id)
				REFERENCES events(id)
				ON DELETE CASCADE,
			FOREIGN KEY (user_id)
				REFERENCES users(id)
				ON DELETE CASCADE
		)
	`
	_, err := DB.Exec(query)

	return err
}

func seedRoles() error {
	query := `
		INSERT INTO roles(id, label) VALUES (?, ?)
	  ON CONFLICT(id) DO UPDATE SET label=excluded.label;
	`

	for _, role := range constant.Roles {
		if _, err := DB.Exec(query, role.Value, role.Label); err != nil {
			return fmt.Errorf("🌱 failed to seed role %q: %w", role.Label, err)
		}
	}

	return nil
}
