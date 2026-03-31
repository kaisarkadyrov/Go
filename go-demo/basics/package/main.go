package main

import (
	employees "demo/app-1/basics/structs"
	"fmt"
)

var x int = 10

func main() {
	x := 5
	fmt.Print(x)
	sayHello("TEST")

	for _, v := range points {
		fmt.Println(v)
	}

	e := employees.Employee{
		"Qaisar",
		18,
		true,
	}

	fmt.Print(e)
}
