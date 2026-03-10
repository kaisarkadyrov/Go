package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}

	fmt.Print(len(slice1))
	fmt.Print(slice1)

	fmt.Print(cap(slice1))

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}
