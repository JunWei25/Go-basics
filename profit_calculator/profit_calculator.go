package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//var revenue float64
	//var expenses float64
	//var taxrate float64

	revenue, err1 := getUserInput("Revenue: ")
	expenses, err2 := getUserInput("Expenses: ")
	taxrate, err3 := getUserInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1) //since the error message will be the same
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxrate)

	fmt.Printf("Earnings Before Tax: %.1f\n", ebt)
	fmt.Printf("Earnings After Tax: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxrate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := float64(earningsBeforeTax) * (1 - taxrate/100)
	ratioTax := earningsBeforeTax / earningsAfterTax
	return earningsBeforeTax, earningsAfterTax, ratioTax
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}
