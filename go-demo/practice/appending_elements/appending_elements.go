package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	numbers := []int{}

	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)

		numbers = append(numbers, number)
	}

	var k int

	fmt.Scan(&k)

	numbers = append(numbers[:k], numbers[k+1:]...)

	fmt.Print(numbers)

}
