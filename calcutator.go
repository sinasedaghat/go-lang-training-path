package main

import (
	"errors"
	"fmt"
)

func main() {
	var revenue, expenses, tax_rate float64
	fmt.Print("enter Revenue: ")
	fmt.Scanln(&revenue)
	err := validate(revenue)
	if err != nil {
		panic(err)
	}
	fmt.Print("enter expenses: ")
	fmt.Scanln(&expenses)
	err = validate(expenses)
	if err != nil {
		panic(err)
	}
	fmt.Print("enter Tax Rate: ")
	fmt.Scanln(&tax_rate)
	err = validate(tax_rate)
	if err != nil {
		panic(err)
	}
	EBT, profit, ratio := calcute(revenue, expenses, tax_rate)
	fmt.Printf("EBT is: %.0f \n", EBT)
	fmt.Printf("Profit is: %.0f \n", profit)
	fmt.Printf("Ratio is: %.2f \n", ratio)
	sign()
}

func sign() {
	fmt.Printf(`######################################
###    ######    #######  ########  ##
##  ###  #### ###  ##### # ###### # ##
###  ########    ####### ## #### ## ##
#####  ###### ###  ##### ### ## ### ##
##  ###  #### ####  #### ####  #### ##
####  ####### ######  ## ########## ##
######################################
`)
}

// func calcute(revenue float64, expenses float64, tax_rate float64) (float64, float64, float64) {
// 	EBT := revenue - expenses
// 	profit := EBT - EBT*(tax_rate/100)
// 	ratio := EBT / profit
// 	return EBT, profit, ratio
// }
func calcute(revenue float64, expenses float64, tax_rate float64) (EBT float64, profit float64, ratio float64) {
	EBT = revenue - expenses
	profit = EBT - EBT*(tax_rate/100)
	ratio = EBT / profit
	return EBT, profit, ratio
}

func validate(input float64) error {
	var err error
	if input <= 0 {
		err = errors.New("input not permitted 0 ")
		return err
	}
	return nil
}
