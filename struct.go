package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName  string
	lastName   string
	birthDate  string
	createDate time.Time
}

func (u user) methodOutputData() {
	fmt.Println("You call method from user struct with methodOutputData() name, Congratulation!", &u)
	fmt.Println("First Name ==> ", u.firstName)
	fmt.Println("Last Name ==> ", u.lastName)
	fmt.Println("Birth Date ==> ", u.birthDate)
	fmt.Println("Create Date ==> ", u.createDate)
}

func (u *user) methodGetArgument(prefix string) {
	fmt.Println("You call method from user struct with methodGetArgument() name, Congratulation!", &u)
	fmt.Println("Argument from methodGetArgument() method in user struct", prefix)
	u.firstName = prefix + " " + u.firstName
}

func main() {
	userFirstName := getUserData("please enter your first name: ")
	userLastName := getUserData("please enter your last name: ")
	userBirthDate := getUserData("please enter your birth date (MM/DD/YYYY): ")

	// var userData2 user
	userData2 := user{
		userFirstName,
		userLastName,
		userBirthDate,
		time.Now(),
	}

	userData := user{
		firstName: userFirstName,
		lastName:  getUserData("please enter your last name (in struct): "),
	}

	userData.methodOutputData()

	userData.methodGetArgument("MR")

	userData.methodOutputData()

	fmt.Println("userData", &userData)

	outputUserData(&userData)
	outputUserData(&userData2)
}

func outputUserData(u *user) {
	fmt.Println("You call outputUserData() function", &u)
	fmt.Println("First Name ==> ", (*u).firstName)
	fmt.Println("Last Name ==> ", u.lastName)
	fmt.Println("Birth Date ==> ", u.birthDate)
	fmt.Println("Create Date ==> ", u.createDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string

	fmt.Scan(&value)
	return value
}
