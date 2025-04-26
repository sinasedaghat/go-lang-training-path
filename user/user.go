package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	// FirstName  string
	firstName  string
	lastName   string
	birthDate  string
	createDate time.Time
}

type Admin struct {
	password string
	email    string
	// Xx       User
	User
}

type Str string

func (txt Str) Log(prefix string) {
	fmt.Println(prefix, txt)
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth day are required")
	}

	return &User{
		firstName:  firstName,
		lastName:   lastName,
		birthDate:  birthDate,
		createDate: time.Now(),
	}, nil
}

func NewAdmin(password, email string) Admin {
	return Admin{
		password: password,
		email:    email,
		// Xx: User{
		User: User{
			firstName:  "Reza",
			lastName:   "Motti",
			birthDate:  "13/2/1372",
			createDate: time.Now(),
		},
	}
}

func (a Admin) MethodOutputAdminDate() {
	// fmt.Println("Admin struct ==> ", a.Xx, a.Xx, a.Xx.birthDate, a.email, a.password)
	fmt.Println("Admin struct ==> ", a.firstName, a.lastName, a.birthDate, a.email, a.password)
}

func (u User) MethodOutputData() {
	// fmt.Println("You call method from user struct with methodOutputData() name", &u)
	fmt.Println("User struct ==> ", u.firstName, u.lastName, u.birthDate)
}

func (u *User) MethodGetArgument(prefix string) {
	// fmt.Println("You call method from user struct with methodGetArgument() name", u) // how to access address of u pointer
	// fmt.Println("Argument from methodGetArgument() method in user struct", prefix)
	u.firstName = prefix + u.firstName
}

func (u *User) MethodClearData() {
	// fmt.Println("You call method from user struct with methodClearData() name", u)
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}
