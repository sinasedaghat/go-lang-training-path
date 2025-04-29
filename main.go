package main

import (
	"fmt"
	"note/fileManagement"
)

func main() {
	title, content := getNodeContent()
	fileManagement.InitSingleton()
	// fmt.Println("practice-getting-user-input ", title, content)
	fileManagement.AddNote(title, content)
}

func getNodeContent() (title, content string) {
	fmt.Print("Note title: ")
	fmt.Scanln(&title)
	fmt.Print("Note content: ")
	fmt.Scan(&content)
	return
}
