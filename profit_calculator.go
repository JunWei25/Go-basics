package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxrate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxrate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := float64(earningsBeforeTax) * (1 - taxrate/100)
	ratioTax := earningsBeforeTax / earningsAfterTax

	fmt.Println(earningsBeforeTax)
	fmt.Println(earningsAfterTax)
	fmt.Println(ratioTax)
}
