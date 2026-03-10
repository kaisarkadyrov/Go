package main

import "fmt"

func main() {

	var n int

	fmt.Scan(n)
	fruits := []int{}

	for i := 1; i < 10; i++ {
		fruits = append(fruits, i)
	}

	fmt.Print(fruits)
}
