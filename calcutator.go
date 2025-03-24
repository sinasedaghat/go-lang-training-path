package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, tax_rate float64
	fmt.Print("enter Revenue: ")
	fmt.Scanln(&revenue)
	fmt.Print("enter expenses: ")
	fmt.Scanln(&expenses)
	fmt.Print("enter Tax Rate: ")
	fmt.Scanln(&tax_rate)
	EBT := revenue - expenses
	profit := EBT - EBT*(tax_rate/100)
	ratio := EBT / profit
	fmt.Printf("EBT is: %f \n", EBT)
	fmt.Printf("Profit is: %f \n", profit)
	fmt.Printf("Ratio is: %f \n", ratio)

}
