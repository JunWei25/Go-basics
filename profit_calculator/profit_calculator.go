package main

import (
	"fmt"
)

func main() {
	//var revenue float64
	//var expenses float64
	//var taxrate float64

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxrate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxrate)

	fmt.Printf("Earnings Before Tax: %.1f\n", ebt)
	fmt.Printf("Earnings After Tax: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxrate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := float64(earningsBeforeTax) * (1 - taxrate/100)
	ratioTax := earningsBeforeTax / earningsAfterTax
	return earningsBeforeTax, earningsAfterTax, ratioTax
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
