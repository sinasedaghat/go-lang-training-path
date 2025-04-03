package main

import (
	"fmt"
	"os"
)

const fileName = "resultCalculated.txt"

func main() {
	// var revenue, expenses, taxRate float64

	// getDirectlyValue("Revenue: ", &revenue)
	// getDirectlyValue("Expenses from getDirectlyValue function: ", &expenses)
	// getDirectlyValue("Tax Rate Percentage: ", &taxRate)

	revenue := getValue("Revenue from getValue function: ", "positive")
	expenses := getValue("Expenses from getDirectlyValue function: ", "positive")
	taxRate := getValue("Tax Rate Percentage: ", "percentage")

	EBT, profit, ratio := profitCalculate(revenue, expenses, taxRate)

	fmt.Printf("Earning Before Tax: %.1f\n", EBT)
	fmt.Printf("Earning After Tax: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func profitCalculate(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = (1 - taxRate/100) * ebt
	ratio = ebt / profit
	writeToFile(ebt, profit, ratio)
	return
}

func getValue(title, valueType string) (value float64) {
	fmt.Print(title)
	fmt.Scan(&value)

	switch valueType {
	case "positive":
		if value <= 0 {
			panic("the value must greater than zero")
		}
	case "percentage":
		if value <= 0 || value >= 100 {
			panic("value is out of range.")
		}
	}
	return value
}

func writeToFile(ebt, profit, ratio float64) {
	data := fmt.Sprintf("Earning Before Tax (ebt): %.1f\nEarning After Tax (profit): %.1f\nratio: %.3f", ebt, profit, ratio)
	os.WriteFile(fileName, []byte(data), 0644)
}

// func getDirectlyValue(title string, variable *float64) {
// 	fmt.Print(title)
// 	fmt.Scan(variable)
// }
