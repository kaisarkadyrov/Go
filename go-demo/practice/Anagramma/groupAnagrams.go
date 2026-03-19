package main

import (
	"fmt"
	"strings"
)

func isAnagram(a, b string) bool {

	a, b = strings.ToLower(a), strings.ToLower(b)

	a, b = strings.TrimSpace(a), strings.TrimSpace(b)

	check1 := map[rune]int{}
	check2 := map[rune]int{}

	runes_a, runes_b := []rune(a), []rune(b)

	for _, s := range runes_a {
		check1[s]++
	}

	for _, s := range runes_b {
		check2[s]++
	}

	if len(check1) != len(check2) {
		return false
	}

	for ind, el := range check1 {
		if check2[ind] != el {
			return false
		}
	}
	return true
}

// func GroupAnagrams(words []string) [][]string {

// }

func main() {

	var a, b string

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Print(isAnagram(a, b))

}
