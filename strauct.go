package main

import "fmt"

func main() {
	firstName := getUserData("please enter your first name: ")
	lastName := getUserData("please enter your last name: ")
	birthDate := getUserData("please enter your birth date (MM/DD/YYYY): ")

	// ... do something awesome whit the gathered data!

	outputUserDate(firstName, lastName, birthDate)
}

func outputUserDate(firstName, lastName, birthDate string) {
	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string

	fmt.Scan(&value)
	return value
}
