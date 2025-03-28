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
	fmt.Printf("EBT is: %.0f \n", EBT)
	fmt.Printf("Profit is: %.0f \n", profit)
	fmt.Printf("Ratio is: %.2f \n", ratio)
	fmt.Printf(`######################################
###    ######    #######  ########  ##
##  ###  #### ###  ##### # ###### # ##
###  ########    ####### ## #### ## ##
#####  ###### ###  ##### ### ## ### ##
##  ###  #### ####  #### ####  #### ##
####  ####### ######  ## ########## ##
######################################`)

}
