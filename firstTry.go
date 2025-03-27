package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 2000
	var years = 10
	var expectedReturnRate = 5.5

	var futureValue = float64(investmentAmount) * math.Pow(1+(expectedReturnRate/100), float64(years))

	fmt.Println(futureValue)
}
