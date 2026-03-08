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

	pairs := map[int]int{}

	for _, num := range numbers {
		pairs[num]++
	}

	for ind, el := range pairs {
		fmt.Printf("%d -> %d \n", ind, el)
	}
}
