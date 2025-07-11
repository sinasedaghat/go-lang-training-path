package prices

import (
	"fmt"

	"url.com/price-calculator/conversion"
	// "url.com/price-calculator/filemanagment"
	"url.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	// filemanagment.FileManagment
	// IOManager         filemanagment.FileManagment `json:"-"`
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	// lines, err := job.InputReader()
	lines, err := job.IOManager.InputReader()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	result := make(map[string]string, len(job.InputPrices))

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	job.TaxIncludedPrices = result

	// err := job.OutputWriter(job)
	err := job.IOManager.OutputWriter(job)

	if err != nil {
		return err
	}
	return nil
}

// func NewTaxIncludedPriceJob(inputPath, outputPath string, taxRate float64) *TaxIncludedPriceJob {
// 	return &TaxIncludedPriceJob{
// 		FileManagment: filemanagment.New(inputPath, outputPath),
// 		TaxRate:       taxRate,
// 	}
// }

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}
