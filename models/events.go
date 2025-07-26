package models

import "time"

var events = []Event{}

type Event struct {
	ID          int       `json:"id"`
	UserId      int       `json:"-"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
	CreateDate  time.Time `json:"createDate"`
	// ID          int       ``
	// UserId      int       ``
	// Name        string    `binding:"required"`
	// Description string    `binding:"required"`
	// Location    string    `binding:"required"`
	// DueDate     time.Time `binding:"required"`
	// CreateDate  time.Time ``
}

func (e Event) Save() {
	// later save in database
	events = append(events, e)
}

func GetEvents() []Event {
	return events
}
