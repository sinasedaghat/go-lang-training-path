package main

import (
	"fmt"
	"regexp"
)

func main() {
	var Email string
	const EmailPattern = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	fmt.Println("Enter Your Email:")
	fmt.Scanln(&Email)
	var state bool
	state, _ = regexp.MatchString(EmailPattern, Email)
	fmt.Println(state)
	if state {
		fmt.Println("Email is correct")
	} else {
		fmt.Println("Email is wrong")
	}
}
