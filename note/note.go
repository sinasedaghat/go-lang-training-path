package note

import (
	"time"

	"github.com/google/uuid"
)

type Node struct {
	// id         uuid.UUID
	// title      string
	// content    string
	// createDate time.Time

	Id         uuid.UUID
	Title      string
	Content    string
	CreateDate time.Time
}

func New(title, content string) *Node {
	return &Node{
		Id:         uuid.New(),
		Title:      title,
		Content:    content,
		CreateDate: time.Now(),
	}
}
