package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate = 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// output information
	// fmt.Println("Future Value:", futureValue)
	//fmt.Printf("Future Value: %v\nFuture Value (adjusted for Inflation): %v", futureValue, futureRealValue)
	formattedFV := fmt.Sprintf("FutureValue: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
	//fmt.Printf(`Future Value: %.1f\n Future Value(adjusted for Inflation): %.1f`,
	//futureValue, futureRealValue)
	//fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for Inflation): %.1f", futureValue, futureRealValue)

}

func outputText(text string) {
	fmt.Print(text)
}

// func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
// 	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv := fv / math.Pow(1+inflationRate/100, years)
// 	return fv, rfv
// }

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}
