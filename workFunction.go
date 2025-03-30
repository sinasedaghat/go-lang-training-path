package main

import (
	"fmt"
	"math"
)

const inflationRate = 3.4

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Years: ")
	fmt.Scan(&years)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := futureValueCalculate(investmentAmount, expectedReturnRate, years)

	futureValueFormatted, futureRealValueFormatted := futureValueFormatted(futureValue, futureRealValue)

	fmt.Println("Future Value: ", futureValue)
	fmt.Println("Future Real Value: ", futureRealValue)

	fmt.Println(futureValueFormatted)
	fmt.Println(futureRealValueFormatted)

	formatOutputText(futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func futureValueCalculate(investmentAmount, expectedReturnRate, years float64) (fv, rfv float64) {
	fv = investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+(inflationRate/100), years)
	// return
	return fv, rfv
}

func futureValueFormatted(fv, rfv float64) (string, string) {
	fvFormatted := fmt.Sprintf("futureValue from futureValueFormatted variable: %v", fv)
	rfvFormatted := fmt.Sprintf("futureRealValue from futureRealValueFormatted variable %.3f", rfv)
	return fvFormatted, rfvFormatted
}

func formatOutputText(fv, rfv float64) {
	fmt.Printf("from Printf ==> \nFuture Value: %v\nFuture Real Value: %v\n", fv, rfv)
	fmt.Printf("from Printf rounded numbers ==> \nFuture Value: %.1f\nFuture Real Value: %.2f\n", fv, rfv)
	fmt.Printf(`from Printf with "Backtick" ==>
	Future Value: %v
	Future Real Value: %v`, fv, rfv)
}
