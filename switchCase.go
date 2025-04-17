package main

import "fmt"

func main() {
	fmt.Println("i'm working on switch case concept")
}

func basicExample(num int) {
	switch num {
	case 1:
		{
			fmt.Println("num is 1")
		}
	case 2:
		{
			fmt.Println("num is 2")
		}
	default:
		{
			fmt.Println("num is'n 1 or 2")
		}
	}
}

func expressionLessExample(num int) {
	switch {
	case num < 0:
		fmt.Println("The number is LESS than 0.")
	case num > 0:
		fmt.Println("The number is GREATER than 0.")
	default:
		fmt.Println("The number is 0.")
	}
}

func fallthroughExample(num int) {
	fmt.Println("If one of the items is true, the commands are executed until 'fallthrough' is not present.")
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three") // This runs because of fallthrough
	default:
		fmt.Println("Other")
	}
}

func multiValuesCaseExample(num int) {
	switch {
	case num < 0, num > 5:
		fmt.Println("The number is out of range.")
	default:
		fmt.Println("The number is in the range.")
	}
}

func stringValueExample(day string) {
	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Printf("%v is weekday.\n", day)
	case "saturday", "sunday":
		fmt.Printf("%v is weekend.\n", day)
	default:
		fmt.Println("Unknown day")
	}
}

func functionCaseExample() {
	switch returnNumber() {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Tow")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Greater Then Five")
	}
}

func returnNumber() int {
	fmt.Println("You call returnNumber() function")
	return 3
}
