package main

import "fmt"

func BirdsInWeek(birdsPerDay []int, week int) int {
	numberOfBirds := map[int]int{}
	var res int
	for i := 0; i < len(birdsPerDay); i++ {
		res += birdsPerDay[i]
		if (i+1)%7 == 0 {
			numberOfBirds[i+1] = res
			res = 0
		}
	}

	weekResult := numberOfBirds[week*7]
	return weekResult
}
func main() {
	t := []int{1, 2, 3}
	fmt.Print(BirdsInWeek([]int{4, 7, 3, 2, 1, 1, 2, 0, 2, 3, 2, 7, 1, 3, 0, 6, 5, 3, 7, 2, 3}, 3))
	fmt.Print(t[3])
}
