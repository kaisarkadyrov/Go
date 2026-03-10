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

	var sum int

	for i := 0; i < n; i++ {
		sum += numbers[i]
	}

	// for better solution we can use range
	// for _, num := range numbers {   here "_" is for index (we do not need it),
	// 		sum += num
	// }

	fmt.Print(sum)

}
