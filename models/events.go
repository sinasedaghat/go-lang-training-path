package models

import (
	"fmt"
	"time"

	"sample-url.com/REST-API/database"
)

type Event struct {
	ID          int       `json:"id"`
	UserId      int       `json:"-"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DueDate     time.Time `json:"due_date" binding:"required"`
	CreateDate  time.Time `json:"create_date"`
	// ID          int       ``
	// UserId      int       ``
	// Name        string    `binding:"required"`
	// Description string    `binding:"required"`
	// Location    string    `binding:"required"`
	// DueDate     time.Time `binding:"required"`
	// CreateDate  time.Time ``
}

func (e Event) Save() error {
	query := `
	INSERT INTO events(user_id, name, description, location, due_date)
	VALUES (?, ?, ?, ?, ?)
	`

	statement, err := database.DB.Prepare(query)
	if err != nil {
		fmt.Println("⚠️ error from database.DB.Prepare() in save function ===>", err)
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(e.UserId, e.Name, e.Description, e.Location, e.DueDate)
	if err != nil {
		fmt.Println("⚠️ error from statement.Exec() in save function ===>", err)
		return err
	}

	fmt.Println("result from save events ===> ", result)

	// index, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Println("⚠️ error from result.LastInsertId() in save function ===>", err)
	// }
	// fmt.Println("index from save event ===> ", index)

	// e.ID = int(index) // update e id from ID database AUTO_INCREMENT

	return err
}

func GetEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := database.DB.Query(query)

	if err != nil {
		fmt.Println("⚠️ error from database.DB.Query() in getEvents function ===>", err)
		return nil, err
	}
	fmt.Println("rows from getEvents ===> ", rows)
	defer rows.Close()

	var events = []Event{}
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.UserId, &event.Name, &event.Description, &event.Location, &event.DueDate, &event.CreateDate)

		if err != nil {
			fmt.Println("⚠️ error from rows.Scan() for loop in getEvents function ===>", err)
			return nil, err
		}
		fmt.Println("event from rows.Scan() for loop in getEvents function ===>", event)

		events = append(events, event)
	}

	return events, nil
}
