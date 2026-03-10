package main

import "fmt"

func contains(set []int, num int) bool {
	for _, v := range set {
		if v == num {
			return false
		}
	}
	return true
}

func main() {

	var n int

	fmt.Scan(&n)

	numbers := []int{}

	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)

		numbers = append(numbers, number)
	}

	set := []int{}

	for _, num := range numbers {
		if contains(set, num) {
			set = append(set, num)
		}
	}

	fmt.Print(set)

}
