package main

import (
	"fmt"

	"example.com/bank/counter"
	"example.com/bank/fileInteraction"
)

const balanceFileName = "balance.txt"

func main() {
	fmt.Printf("This Variable is newMainVariable = %v\n This Print from bank.go file (main package.)", newMainVariable)

	var choice int
	balance, err := fileInteraction.GetNumberFromFile(balanceFileName)

	if err != nil {
		fmt.Println("❌ The bank is out of reach.")
		fmt.Println("❌", err)
		return
		// msg := fmt.Sprintf("❌ The bank is out of reach.\n%v", err)
		// panic(msg)
	}

	fmt.Println("Welcome to Bank 🏦")

	for { // for i := 0; i < 5; i++ { // for range 5 {
		counter.Increase()
		selectTypeService()

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
			fileInteraction.WriteNumberToFile(balanceFileName, balance)
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
			fileInteraction.WriteNumberToFile(balanceFileName, balance)
		} else if choice == 4 {
			fmt.Println("Goodby! 👋")
			break
		} else if choice == 5 {
			fmt.Printf("Loop counter: %v\n", counter.Counter)
		} else {
			fmt.Println("Your choice isn't valid! 😞")
		}
	}

	fmt.Printf("📉 final counter %v\n", counter.Counter)
	println("🏦 Thanks for choosing our bank.")
}
