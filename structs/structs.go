package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ") //cant handle spaces
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")

	//when using anonymous embedding
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// when using explicit embedding
	// admin.User.OutputUserDetails()
	// admin.User.ClearUserName()
	// admin.User.OutputUserDetails()

	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	//Scan will wait till user enters a value even if the user press enter.
	fmt.Scanln(&value)
	return value
}
