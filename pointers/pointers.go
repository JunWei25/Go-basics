package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	agePointer = &age

	fmt.Println("Age:", agePointer)
	fmt.Println("Age:", *agePointer)

	fmt.Println("Age: ", age)
	editAgeToAdultYears(agePointer)
	//adultYears := getAdultYears(agePointer)
	//fmt.Println(adultYears)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18

}
