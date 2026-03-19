package main

import "fmt"

func Reverse(s string) string {
	runes := []rune(s) // для работы с кириллицей и эмодзи

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {

	var n string

	fmt.Scan(&n)

	fmt.Print(Reverse(n))
}
