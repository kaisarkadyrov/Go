package main

import (
	"fmt"
)

func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func returnFunc(x string) func() {
	sum := 0
	return func() {
		fmt.Println(sum)
	}
}

func activateGiftCard() func(int) int {
	amount := 100

	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}

	return debitFunc
}

func main() {

	test := func(x int) int {
		return x * -1
	}

	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()
	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard2(5))

	test2(test)
	returnFunc("hello")()

}
