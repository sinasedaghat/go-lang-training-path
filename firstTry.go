package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 3.4
	var investmentAmount float64 = 2000
	years := 10.0             // var years float64 = 10
	expectedReturnRate := 5.5 // var expectedReturnRate = 5.5

	futureValue := investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+(inflationRate/100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
