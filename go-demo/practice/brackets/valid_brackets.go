package main

import "fmt"

// test commit for activity
func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (ch == ')' && top != '(') ||
				(ch == ']' && top != '[') ||
				(ch == '}' && top != '{') {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("([]){}"))

}
