package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrom(s string) bool {

	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, _ := reader.ReadString('\n')

	if isPalindrom(n) {
		fmt.Print("Palindrom")
	} else {
		fmt.Print("Not Palindrom")
	}
}
