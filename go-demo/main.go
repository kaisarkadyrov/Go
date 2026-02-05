package main

import (
	"fmt"
	"math"
)

func main() {
	// var userHeight, userWeight float64 = 1.8, 100

	// var userHeight = 1.8
	// var userWeight float64 = 100

	userHeight := 1.8
	userWeight := 100.0

	// var IMT = userWeight / math.Pow(userHeight, 2)
	IMT := userWeight / math.Pow(userHeight, 2)

	fmt.Print(IMT)

}
