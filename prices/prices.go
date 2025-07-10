package prices

import (
	"fmt"

	"url.com/price-calculator/conversion"
	"url.com/price-calculator/filemanagment"
)

type TaxIncludedPriceJob struct {
	// filemanagment.FileManagment
	FileManagment     filemanagment.FileManagment `json:"-"`
	TaxRate           float64                     `json:"tax_rate"`
	InputPrices       []float64                   `json:"input_prices"`
	TaxIncludedPrices map[string]string           `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	// lines, err := job.LinesReader()
	lines, err := job.FileManagment.LinesReader()

	if err != nil {
		fmt.Println("Can't read file!")
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
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

	job.TaxIncludedPrices = result

	// err := job.WriteJSON(job)
	err := job.FileManagment.WriteJSON(job)

	if err != nil {
		fmt.Println("Can't write in file")
		fmt.Println(err)
		return
	}
}

func NewTaxIncludedPriceJob(inputPath, outputPath string, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		FileManagment: filemanagment.New(inputPath, outputPath),
		TaxRate:       taxRate,
	}
}
