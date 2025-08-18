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

func GetEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := database.DB.Query(query)

	if err != nil {
		fmt.Println("🗂️ Error from Query(query) of GetEvents function:", err)
		return nil, err
	}
	defer rows.Close()

	var events = []Event{}
	for rows.Next() {
		var event Event
		// err = rows.Scan(&event.ID, &event.UserId, &event.Name, &event.Description, &event.Location, &event.DueDate, &event.CreateDate)

		err = rows.Scan(
			&event.ID,
			&event.UserId,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DueDate,
			&event.CreateDate,
		)

		if err != nil {
			fmt.Println("🗂️ Error from Scan(pointers...) of GetEvents function:", err)
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEvent(id int) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := database.DB.QueryRow(query, id)

	var event Event
	// err := row.Scan(&event.ID, &event.UserId, &event.Name, &event.Description, &event.Location, &event.DueDate, &event.CreateDate)

	err := row.Scan(
		&event.ID,
		&event.UserId,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DueDate,
		&event.CreateDate,
	)

	if err != nil {
		fmt.Println("🗂️ Error from Scan(pointers...) of GetEvent function:", err)
		return nil, err
	}

	return &event, nil
}

func (e Event) Save() error {
	query := `
	INSERT INTO events(user_id, name, description, location, due_date)
	VALUES (?, ?, ?, ?, ?)
	`

	statement, err := database.DB.Prepare(query)
	if err != nil {
		fmt.Println("🗂️ Error from Prepare(query) of event.Save method:", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.UserId, e.Name, e.Description, e.Location, e.DueDate)
	// result, err := statement.Exec(e.UserId, e.Name, e.Description, e.Location, e.DueDate)
	if err != nil {
		fmt.Println("🗂️ Error from Exec(data...) of event.Save method:", err)
	}

	// index, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Println("⚠️ error from result.LastInsertId() in save function ===>", err)
	// }
	// fmt.Println("index from save event ===> ", index)

	// e.ID = int(index) // update e id from ID database AUTO_INCREMENT

	return err
}

func (e Event) Update() error {
	fmt.Println("event from model.update", e)
	query := `
		UPDATE events
		SET user_id = ?, name = ?, description = ?, location = ?, due_date = ?
		WHERE id = ?; 
	`
	statement, err := database.DB.Prepare(query)
	if err != nil {
		fmt.Println("🗂️ Error from Prepare(query) of event.Update method:", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.UserId, e.Name, e.Description, e.Location, e.DueDate, e.ID)
	if err != nil {
		fmt.Println("🗂️ Error from Exec(data...) of event.Update method:", err)
	}

	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	statement, err := database.DB.Prepare(query)
	if err != nil {
		fmt.Println("🗂️ Error from Prepare(query) of event.Delete method:", err)
		return err
	}

	_, err = statement.Exec(e.ID)
	if err != nil {
		fmt.Println("🗂️ Error from Exec(data...) of event.Delete method:", err)
	}

	return err
}

func (e Event) Register(userId int) error {
	query := "INSERT INTO registration(event_id, user_id) VALUES (?, ?)"

	statement, err := database.DB.Prepare(query)

	if err != nil {
		fmt.Println("🗂️ Error from Prepare(query) of event.Register method:", err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.ID, userId)

	if err != nil {
		fmt.Println("🗂️ Error from Exec(data...) of event.Register method:", err)
	}
	return err
}

func (e Event) Unregister(userId int) error {
	query := "DELETE FROM registration WHERE event_id = ? AND user_id = ?"

	statement, err := database.DB.Prepare(query)

	if err != nil {
		fmt.Println("🗂️ Error from Prepare(query) of event.Unregister method:", err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.ID, userId)

	if err != nil {
		fmt.Println("🗂️ Error from Exec(data...) of event.Unregister method:", err)
	}
	return err
}
