package main

import (
	"fmt"

	// "url.com/price-calculator/cmdmanager"
	"url.com/price-calculator/filemanagment"
	"url.com/price-calculator/prices"
)

var taxRates []float64 = []float64{}

func main() {
	taxRates = []float64{0, 0.07, 0.1, 0.15}

	for _, rate := range taxRates {
		// priceJob := prices.NewTaxIncludedPriceJob("prices.txt", fmt.Sprintf("tax_%v.json", rate), rate)
		fm := filemanagment.New("prices.txt", fmt.Sprintf("tax_%v.json", rate))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, rate)

		err := priceJob.LoadData()
		if err != nil {
			fmt.Println("has error")
			fmt.Println(err)
		}
		err = priceJob.Process()
		if err != nil {
			fmt.Println("has error")
			fmt.Println(err)
		}
	}
}
