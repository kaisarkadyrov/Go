package main

import "fmt"

func main() {

	m := map[string]int{"apple": 5, "orange": 10}

	value, ok := m["apple"]

	if ok {
		fmt.Print("Founded: ", value)
	} else {
		fmt.Print("Not found")
	}

	value, ok = m["cherry"]

	if ok {
		fmt.Print("Founded: ", value)
	} else {
		fmt.Print("Not Founded")
	}
	fmt.Print(ok)

}
