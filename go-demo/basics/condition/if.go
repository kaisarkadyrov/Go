package main

import "fmt"

func main() {

	var a int

	fmt.Scan(&a)

	if a > 10 {
		fmt.Print("The number is more than 10")
	} else if a < 10 {
		fmt.Print("The number is less than 10")
	} else {
		fmt.Print("The number is 10")
	}

}
