package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 2000
	years := 10.0             // var years float64 = 10
	expectedReturnRate := 5.5 // var expectedReturnRate = 5.5

	var futureValue = investmentAmount * math.Pow(1+(expectedReturnRate/100), years)

	fmt.Println(futureValue)
}
