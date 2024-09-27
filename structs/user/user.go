package user

import (
	"errors"
	"fmt"
	"time"
)

// need to user uppercase in order to export, also applies to functions and fields
// use lowercase if dont want to expose fields/methods/functions
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// func outputUserDetails(u *user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

// Method with value receiver
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// Mutation methods, using a pointer to edit value
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name, birthdate are required.")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
