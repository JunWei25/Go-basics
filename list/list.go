package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	prices[2] = 11.99

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices)

	prices = append(prices, 5.99)
	fmt.Println(prices)
}

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	productNames[2] = "A Carpet"

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
