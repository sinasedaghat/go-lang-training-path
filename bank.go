package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func main() {
	var counter int
	var choice int
	balance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("❌ The bank is out of reach.")
		fmt.Println("❌", err)
		return
		// msg := fmt.Sprintf("❌ The bank is out of reach.\n%v", err)
		// panic(msg)
	}

	fmt.Println("Welcome to Bank 🏦")

	for { // for i := 0; i < 5; i++ { // for range 5 {
		counter++
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. 💰 Check balance")
		fmt.Println("2. 💵 Deposit money")
		fmt.Println("3. 💸 Withdraw money")
		fmt.Println("4. 👋 Exit")
		fmt.Println("5. Print loop counter")

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
			writeBalanceToFile(balance)
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
			writeBalanceToFile(balance)
		} else if choice == 4 {
			fmt.Println("Goodby! 👋")
			break
		} else if choice == 5 {
			fmt.Printf("Loop counter: %v\n", counter)
		} else {
			fmt.Println("Your choice isn't valid! 😞")
		}

		fmt.Printf("📈 counter %v\n", counter)
	}

	fmt.Printf("📉 final counter %v\n", counter)
	println("🏦 Thanks for choosing our bank.")
}

func writeBalanceToFile(balance float64) {
	os.WriteFile(balanceFileName, []byte(fmt.Sprint(balance)), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFileName)
	fmt.Println("🔮 This is ERROR received when i want get balance data from file.", err)
	if err != nil {
		return 0, errors.New("your Balance data does not exists")
	}

	balance, err := strconv.ParseFloat(string(data), 64)
	fmt.Println("🔮 This is ERROR received when i want convert string from file to floating point number.", err)
	if err != nil {
		return 0, errors.New("your Balance isn't valid value")
	}
	return balance, nil
}
