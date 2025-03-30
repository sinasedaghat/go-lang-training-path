package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	// getDirectlyValue("Revenue: ", &revenue)
	revenue = getValue("Revenue from getValue function: ")
	getDirectlyValue("Expenses from getDirectlyValue function: ", &expenses)
	getDirectlyValue("Tax Rate Percentage: ", &taxRate)

	EBT, profit, ratio := profitCalculate(revenue, expenses, taxRate)

	fmt.Printf("Earning Before Tax: %.1f\n", EBT)
	fmt.Printf("Earning After Tax: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
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
