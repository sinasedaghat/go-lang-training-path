package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)
	// getDirectlyValue("Revenue: ", &revenue)
	revenue = getValue("Revenue from getValue function: ")

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)
	getDirectlyValue("Expenses from getDirectlyValue function: ", &expenses)

	// fmt.Print("Tax Rate Percentage: ")
	// fmt.Scan(&taxRate)
	getDirectlyValue("Tax Rate Percentage: ", &taxRate)

	// EBT := revenue - expenses
	// profit := (1 - taxRate/100) * EBT
	// ratio := EBT / profit
	EBT, profit, ratio := profitCalculate(revenue, expenses, taxRate)

	// fmt.Print("Earning Before Tax: ")
	// fmt.Println(EBT)
	fmt.Println("Earning Before Tax: ", EBT)

	// fmt.Print("Earning After Tax: ")
	// fmt.Println((profit))
	fmt.Println("Earning After Tax: ", profit)

	// fmt.Print("Ratio: ")
	// fmt.Println(ratio)
	fmt.Println("Ratio: ", ratio)
}

func profitCalculate(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = (1 - taxRate/100) * ebt
	ratio = ebt / profit

	return
}

func getValue(title string) (value float64) {
	fmt.Print(title)
	fmt.Scan(&value)
	return value
}

func getDirectlyValue(title string, variable *float64) {
	fmt.Print(title)
	fmt.Scan(variable)
}
