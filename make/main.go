package main

import "fmt"

//useful if know in advance how much items is going to be added
func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	coursesRatings := make(map[string]float64, 3)

	coursesRatings["go"] = 4.7
	coursesRatings["react"] = 4.8
	coursesRatings["angular"] = 4.7

	fmt.Println(coursesRatings)
}
