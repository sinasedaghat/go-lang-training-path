package fileManagement

import (
	"encoding/json"
	"fmt"
	"os"

	"note/note"
)

// const filePath = "myNote/sunday.json"
const fileName = "note.json"

var data []note.Node

func Init() {
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		createFile()
	} else {
		readFile()
	}
}

func createFile() {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func readFile() {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	if len(content) != 0 {
		err = json.Unmarshal(content, &data)
		if err != nil {
			panic(err)
		}
	}
}

func writeFile() {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fileName, content, 0644)
	if err != nil {
		panic(err)
	}
}

func AddNote(title, content string) {
	data = append(data, *note.New(title, content))

	fmt.Println("data readFile ==> ", data)
	writeFile()
}
