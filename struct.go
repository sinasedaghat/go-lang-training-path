package main

import (
	"example/structs/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("please enter your first name: ")
	userLastName := getUserData("please enter your last name: ")
	userBirthDate := getUserData("please enter your birth date (MM/DD/YYYY): ")

	var userData *user.User

	// userData = &user.User{
	// 	FirstName: userFirstName,
	// }

	// name := "Reza"
	// last := "Motti"
	// birth := "13/2/1372"

	// userData1, _ := user.New(&name, &last, &birth)
	userData, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	userData.MethodGetArgument("Mr.")
	userData.MethodOutputData()
	userData.MethodClearData()
	userData.MethodOutputData()

	fmt.Printf("first userDate: %v\ntype of userDate: %T\n", userData, userData)
	// fmt.Printf("first userDate1: %v\ntype of userDate: %T\n", userData1, userData1)
	// fmt.Printf("second userDate: %v\ntype of userDate: %T\n", userData, userData)
	// fmt.Printf("type of FirstName %T\n", &userFirstName)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string

	fmt.Scanln(&value)
	return value
}
