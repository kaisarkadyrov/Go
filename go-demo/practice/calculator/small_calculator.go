package main

import "fmt"

func main() {
	var a int
	var b int

	var oper string

	fmt.Scan(&a, &b, &oper)

	switch oper {
	case "+":
		fmt.Printf("The sum of %d and %d is %d ", a, b, a+b)
	case "-":
		fmt.Printf("The substraction of %d and %d is %d ", a, b, a-b)
	case "*":
		fmt.Printf("The multiplication of %d and %d is %d ", a, b, a*b)
	case "/":
		if b == 0 {
			fmt.Print("Division by 0 is not possible")
		} else {
			var a_for_div float64 = float64(a)
			var b_for_div float64 = float64(b)
			fmt.Printf("The division of %d and %d is %.1f ", a, b, a_for_div/b_for_div) //not perfect but ok
		}
	default:
		fmt.Println("Unknown operation")
	}

}
