package main

import "fmt"

func main() {
	var choice int
	balance := 1000.0

	fmt.Println("Welcome to Bank 🏦")
	fmt.Println("What do you want to do?")
	fmt.Println("1. 💰 Check balance")
	fmt.Println("2. 💵 Deposit money")
	fmt.Println("3. 💸 Withdraw money")
	fmt.Println("4. 👋 Exit")

	fmt.Printf("\n❓ Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		// println("Your balance is:", balance) // why this code is work? // this print in my terminal ==> Your balance is: +1.000000e+003
		fmt.Println("💰 Your balance is:", balance)
	} else if choice == 2 {
		var input float64
		fmt.Print("💵 Your deposit: ")
		fmt.Scan(&input)
		balance += input
		fmt.Printf("💰 Your Balance is Update! You have %.2f\n", balance)
	} else if choice == 3 {
		var input float64
		fmt.Print("💸 Withdrawal amount: ")
		fmt.Scan(&input)
		balance -= input
		fmt.Printf("💰 Your Balance is Update! You have %.2f\n", balance)
	} else if choice == 4 {
		fmt.Println("Goodby! 👋")
	} else {
		fmt.Println("Your choice isn't valid! 😞")
		fmt.Println("Goodby! 👋")
	}
}
