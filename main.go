package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"note/note"
	"note/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	Show()
	saver
}

func main() {
	userNote, err := note.New(getNoteData())
	if err != nil {
		fmt.Println(err)
		return
	}

	err = output(userNote)
	if err != nil {
		return
	}

	userTodo, err := todo.New(getUserInput("Todo text"))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = output(userTodo)
	if err != nil {
		return
	}
}

func output(data outputable) error {
	data.Show()
	return outputSave(data)
	// return data.Save().Error()
}

func outputSave(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}
	fmt.Println("Saving successful.")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v: ", prompt)
	var value string

	value, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}

func getNoteData() (string, string) {
	title := getUserInput("Note title")
	content := getUserInput("Note content")

	return title, content
}
