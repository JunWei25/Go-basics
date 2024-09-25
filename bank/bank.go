package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalancefile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalancefile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		return
		//panic(err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalancefile)
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalancefile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is", accountBalance)
		// 	continue
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}
		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Withdraw amount: ")
		// 	var withdrawalAmount float64
		// 	fmt.Scan(&withdrawalAmount)

		// 	if withdrawalAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}

		// 	if withdrawalAmount > accountBalance {
		// 		fmt.Println("Invalid amount. You can't withdraw more than you have.")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawalAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	// return
		// 	break
		// }
	}
}
