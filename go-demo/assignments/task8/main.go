package main

import "fmt"

func main() {
	var str string
	var letter string

	fmt.Scan(&str)
	fmt.Scan(&letter)

	count := 0

	for i := 0; i < len(str); i++ {
		if string(str[i]) == letter {
			count = count + 1
		}
	}

	fmt.Print(count)
}
