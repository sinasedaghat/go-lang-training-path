package prices

import (
	"fmt"

	"url.com/price-calculator/conversion"
	"url.com/price-calculator/filemanagment"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanagment.LinesReader("prices.txt")

	if err != nil {
		fmt.Println("Cant read file!")
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println("Cant convert file content to prices!")
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]string, len(job.InputPrices))

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		// InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}
