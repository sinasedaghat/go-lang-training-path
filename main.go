package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"note/note"
)

func main() {
	userNote, err := note.New(getNoteData())
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.ShowNote()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving failed.")
		return
	}
	fmt.Println("Saving successful.")
}

func getUserInput(prompt string) string {
	fmt.Printf("%v: ", prompt)
	var value string

	value, err := bufio.NewReader(os.Stdin).ReadString('\n')
	// reader.ReadString('\n')

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
