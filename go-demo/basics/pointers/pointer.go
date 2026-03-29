package main

import "fmt"

type User struct {
	email    string
	username string
	age      int
}

func (u User) Email() string {
	return u.email
}

func (u User) UserName() string {
	return u.username
}

func (u User) Age() int {
	return u.age
}

func main() {

	user := User{
		email:    "stm@foo.com",
		username: "Alex",
		age:      25,
	}

	fmt.Print(user.Email())
	fmt.Print(user.UserName())
	fmt.Print(user.Age())
}
