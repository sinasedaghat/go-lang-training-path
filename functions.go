package main

import "fmt"

type function func(int) int

func main() {
	normalNumbers := []int{1, 2, 3, 4}
	doubleNumbers := transformNumber(&normalNumbers, double)
	tripleNumbers := transformNumber(&normalNumbers, triple)

	fmt.Println("normalNumbers ==>", normalNumbers)
	fmt.Println("doubleNumbers ==>", doubleNumbers)
	fmt.Println("tripleNumbers ==>", tripleNumbers)

	equalityDoubleNumbers := transformNumber(&doubleNumbers, selectActionForEqual(&doubleNumbers))
	equalityTripleNumbers := transformNumber(&tripleNumbers, selectActionForEqual(&tripleNumbers))

	fmt.Println("equalityDoubleNumbers ==>", equalityDoubleNumbers)
	fmt.Println("equalityTripleNumbers ==>", equalityTripleNumbers)
}

// func transformNumber(numbers *[]int, function func(int) int) []int {
func transformNumber(numbers *[]int, function function) []int {
	new := make([]int, 0)

	for _, value := range *numbers {
		new = append(new, function(value))
	}
	return new
}

func selectActionForEqual(numbers *[]int) function {
	if (*numbers)[0]%2 == 0 {
		return triple
	} else {
		return double
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
