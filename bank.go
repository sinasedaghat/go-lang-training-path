package main

import "fmt"

func main() {
	var choice int
	balance := 1000.0
	fmt.Println("Welcome to Bank 🏦")

	for i := 0; i < 5; i++ {
		fmt.Println("What do you want to do?")
		fmt.Println("1. 💰 Check balance")
		fmt.Println("2. 💵 Deposit money")
		fmt.Println("3. 💸 Withdraw money")
		fmt.Println("4. 👋 Exit")

		fmt.Printf("\n❓ Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			// println("Your balance is:", balance) // Why does this code work? // this print in my terminal ==> Your balance is: +1.000000e+003
			fmt.Printf("💰 Your balance is: %.2f\n", balance)
		} else if choice == 2 {
			var input float64
			fmt.Print("💵 Your deposit: ")
			fmt.Scan(&input)
			if input <= 0 {
				fmt.Println("‼️ Invalid amount. Must be greater than 0.00!")
				continue
			}
			balance += input
			fmt.Printf("💰 Your Balance is Update! You have %.2f\n", balance)
		} else if choice == 3 {
			var input float64
			fmt.Print("💸 Withdrawal amount: ")
			fmt.Scan(&input)
			if input <= 0 {
				fmt.Println("‼️ Invalid amount. Must be greater than 0.00!")
				continue
			}
			if input >= balance {
				fmt.Printf("‼️ Invalid amount. Must be lower than %.2f! (your balance)\n", balance)
				continue
			}
			balance -= input
			fmt.Printf("💰 Your Balance is Update! You have %.2f\n", balance)
		} else if choice == 4 {
			fmt.Println("Goodby! 👋")
			break
		} else {
			fmt.Println("Your choice isn't valid! 😞")
		}
	}

	println("🏦 Thanks for choosing our bank.")
}
