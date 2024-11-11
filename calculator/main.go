package main

import (
	"fmt"

	"example.com/calculator/cmdmanager"
	"example.com/calculator/prices"
)

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()

		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}