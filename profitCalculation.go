package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate Percentage: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	profit := (1 - taxRate/100) * EBT
	ratio := EBT / profit

	fmt.Print("Earning Before Tax: ")
	fmt.Println(EBT)

	fmt.Print("Earning After Tax: ")
	fmt.Println((profit))

	fmt.Print("Ratio: ")
	fmt.Println(ratio)
}
