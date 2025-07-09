package main

import "price-calculator/prices"

// var prices []float64 = []float64{}
var taxRates []float64 = []float64{}

// var result map[float64][]float64 = map[float64][]float64{}

func main() {
	// prices = []float64{10, 20, 30}
	taxRates = []float64{0, 0.07, 0.1, 0.15}

	for _, rate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(rate)
		// pricesIncludeTax := make([]float64, len(prices))
		priceJob.Process()

		// for index, price := range prices {
		// 	pricesIncludeTax[index] = price * (1 + rate)
		// }

		// result[rate] = pricesIncludeTax
	}

	// fmt.Println(result)
}
