package main

import (
	"fmt"
	"strings"
)

type Data struct {
	name  string
	email string
	age   int
}

func main() {

	var user Data

	fmt.Scan(&user.name)
	fmt.Scan(&user.email)
	fmt.Scan(&user.age)

	var isName, isEmail, isAge bool = false, false, false

	// checking if the data are correct
	if len(user.name) >= 2 && len(user.name) <= 14 {
		isName = true
	}
	parts := strings.Split(user.email, "@")
	local, domain := parts[0], parts[1]
	if (len(local) >= 4 && len(local) <= 15) && strings.Contains(domain, ".") {
		isEmail = true
	}
	if user.age >= 6 && user.age <= 100 {
		isAge = true
	}

	// control checking
	if !isName {
		fmt.Print("Your name length must be more than 2 and no more than 14 characters")
	}
	if !isEmail {
		fmt.Print("Your email is incorrect")
	}
	if !isAge {
		fmt.Print("Enter your correct age")
	}

	fmt.Print("Succesful!")
}
