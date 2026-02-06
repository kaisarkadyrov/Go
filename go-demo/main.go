package main

import (
	"fmt"
	"math"
)

func main() {
	// var userHeight, userWeight float64 = 1.8, 100

	// var userHeight = 1.8
	// var userWeight float64 = 100
	// var lol int = 50
	var userHeight float64
	var userWeight float64

	// var a, b, c, d int = 1, 3, 5, 7

	fmt.Print("Enter your height: ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight: ")
	fmt.Scan(&userWeight)

	const IMT_power = 2
	// var IMT = userWeight / math.Pow(userHeight, 2)
	IMT := userWeight / math.Pow(userHeight, IMT_power)

	fmt.Printf("Your IMT: %.0f", IMT) // выводит форматированную строку, выводит точное число

	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)

	fmt.Printf("type of the i: %T", i)
}
