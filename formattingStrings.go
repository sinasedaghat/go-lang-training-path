package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 3.4
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+(inflationRate/100), years)

	futureValueFormatted := fmt.Sprintf("futureValue from futureValueFormatted variable: %v", futureValue)
	futureRealValueFormatted := fmt.Sprintf("futureRealValue from futureRealValueFormatted variable %.3f", futureRealValue)

	fmt.Println("Future Value: ", futureValue)
	fmt.Println("Future Real Value: ", futureRealValue)

	fmt.Println(futureValueFormatted)
	fmt.Println(futureRealValueFormatted)

	fmt.Printf("from Printf ==> \nFuture Value: %v\nFuture Real Value: %v\n", futureValue, futureRealValue)
	fmt.Printf("from Printf rounded numbers ==> \nFuture Value: %.1f\nFuture Real Value: %.2f\n", futureValue, futureRealValue)
	fmt.Printf(`from Printf with "Backtick" ==> 
	Future Value: %v
	Future Real Value: %v`, futureValue, futureRealValue)
}
