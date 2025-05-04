package note

import (
	"encoding/json"
	"errors"
	"fmt"

	"os"
	"strings"
	"time"
)

type Note struct {
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"create"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:      title,
		Content:    content,
		CreateDate: time.Now(),
	}, nil
}

func (note Note) ShowNote() {
	fmt.Printf("Your note has %v as title and it's content is %v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
