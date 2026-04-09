package main

import "fmt"

func PrependItems(slice []int, values ...int) ([]int, []int) {
	for _, v := range values {
		slice = append([]int{v}, slice...)
	}
	return slice, values
}

func main() {

	fmt.Print(PrependItems([]int{5, 2, 10, 6, 8, 7, 0, 9}, 0, 6))
}
