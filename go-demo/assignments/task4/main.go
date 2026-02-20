package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	var result int

	for i := 1; i <= n; i++ {
		result = result + i
	}

	fmt.Print(result)
}
