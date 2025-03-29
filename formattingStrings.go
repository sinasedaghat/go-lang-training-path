package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 3.4
	var investmentAmount float64 // var investmentAmount float64 = 2000
	var years float64            // years := 10.0             // var years float64 = 10
	expectedReturnRate := 5.5    // var expectedReturnRate = 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+(inflationRate/100), years)

	fmt.Print("Future Value: ")
	fmt.Println(futureValue)
	fmt.Print("Future Real Value: ")
	fmt.Println(futureRealValue)
}
