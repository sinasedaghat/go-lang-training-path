package main

import (
	"errors"
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
	// fmt.Println("You call method from user struct with methodOutputData() name", &u)
	fmt.Println("user struct ==> ", u)
}

func (u *user) methodGetArgument(prefix string) {
	// fmt.Println("You call method from user struct with methodGetArgument() name", u) // how to access address of u pointer
	// fmt.Println("Argument from methodGetArgument() method in user struct", prefix)
	u.firstName = prefix + " " + u.firstName
}

func (u *user) methodClearData() {
	// fmt.Println("You call method from user struct with methodClearData() name", u)
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

func newUser(firstName, lastName, birthDate *string) (*user, error) {
	if *firstName == "" || *lastName == "" || *birthDate == "" {
		return nil, errors.New("first name, last name and birth day are required")
	}

	return &user{
		firstName:  *firstName,
		lastName:   *lastName,
		birthDate:  *birthDate,
		createDate: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("please enter your first name: ")
	userLastName := getUserData("please enter your last name: ")
	userBirthDate := getUserData("please enter your birth date (MM/DD/YYYY): ")

	// userData := user{
	// 	firstName: userFirstName,
	// 	// lastName:   getUserData("please enter your last name (in struct): "),
	// 	lastName:   userLastName,
	// 	birthDate:  userBirthDate,
	// 	createDate: time.Now(),
	// }
	var userData *user

	userData, err := newUser(&userFirstName, &userLastName, &userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("type of userDate %T\n", userData)
	fmt.Printf("type of FirstName %T\n", &userFirstName)

	userData.methodOutputData()

	// userData.methodGetArgument("MR")
	userData.methodClearData()

	userData.methodOutputData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string

	fmt.Scanln(&value)
	return value
}
