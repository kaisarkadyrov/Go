package main

import "fmt"

func main() {

	// #1
	// fruits := [3]string{"apple", "orange", "banana"}

	// for index, value := range fruits {
	// 	fmt.Printf("Index: %d and Value: %v and type: %T", index, value, value, "\n")
	// }

	// #2
	var n int
	fruits := []string{}

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var fruit string
		fmt.Scan(&fruit)

		fruits = append(fruits, fruit)
	}

	for index, value := range fruits {
		fmt.Printf("Index: %d | Value: %v | Type: %T", index, value, value, "\n")
	}

}
