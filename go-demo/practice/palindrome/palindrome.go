package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(isPalindrome(x))

	number := 4.3242
	r := fmt.Sprintf("%.2f", number)
	fmt.Print(r)
}
