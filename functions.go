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

	// anonymous function
	plusFourArray := transformNumber(&normalNumbers, func(number int) int {
		return number + 4
	})
	fmt.Println("plusFourArray ==>", plusFourArray)
	fmt.Println("normalNumbers ==>", normalNumbers)

	// use closer concept
	double := transformCreator(2)
	triple := transformCreator(3)

	doubled := transformNumber(&normalNumbers, double)
	tripled := transformNumber(&normalNumbers, triple)

	fmt.Println("doubled ==>", doubled)
	fmt.Println("tripled ==>", tripled)

	// recursion
	fibonacciResult := fibonacciRecursion(3)
	fmt.Println("fibonacciResult ==>", fibonacciResult)
	// fmt.Println("fibonacciResult ==>", fibonacciRecursion(0), fibonacciRecursion(1), fibonacciRecursion(2), fibonacciRecursion(3), fibonacciRecursion(4), fibonacciRecursion(5), fibonacciRecursion(6), fibonacciRecursion(7), fibonacciRecursion(8), fibonacciRecursion(9))

	// aggregate has collect all parameter
	totalSum := aggregate(1, 2, 3, 4, 5)
	fmt.Println("totalSum ==>", totalSum)

	//
	totalNormalNumbers := aggregate(normalNumbers...)
	fmt.Println("totalNormalNumbers ==>", totalNormalNumbers)
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

func transformCreator(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// recursion function
func fibonacciRecursion(number int) int {
	if number <= 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		return fibonacciRecursion(number-1) + fibonacciRecursion(number-2)
	}
}

// variadic parameter functions
func aggregate(numbers ...int) (sum int) {
	// sum := 0
	for _, number := range numbers {
		sum += number
	}
	return
}
