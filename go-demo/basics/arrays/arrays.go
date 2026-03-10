package main

import "fmt"

func main() {
	var arr = [4]int{1, 2, 3, 4}

	arr2 := [5]int{10, 20, 30, 40, 50}

	arr2[3] = 100

	arr3 := [5]int{1, 2}
	arr4 := [5]int{}

	//specific elements

	arr5 := [5]int{1: 10, 2: 20, 4: 40}

	fmt.Print(arr)
	fmt.Print(arr2) // 10 20 30 100 50
	fmt.Print(arr3) // 1 2 0 0 0
	fmt.Print(arr4) // 0 0 0 0 0
	fmt.Print(arr5) // 0 10 20 0 40

	arr6 := [...]string{"Volvo", "BMW", "Ford", "Mazda"}

	fmt.Print(len(arr6))
	fmt.Print(len(arr2)) //5

}
