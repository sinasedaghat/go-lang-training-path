package note

import (
	"time"

	"github.com/google/uuid"
)

type Node struct {
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
