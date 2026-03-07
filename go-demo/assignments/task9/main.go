package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')

	count := 0

	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			count++
		}
	}

	fmt.Print(count + 1)
}
