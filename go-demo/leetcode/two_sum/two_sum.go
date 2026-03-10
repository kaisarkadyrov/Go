package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	nums := []int{}

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')

	parts := strings.Fields(line)

	for _, p := range parts {
		num, _ := strconv.Atoi(p)
		nums = append(nums, num)
	}

	var target int
	fmt.Scan(&target)

	var res_i int
	var res_j int

	for i := 0; i < len(nums); i++ {

		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res_i = i
				res_j = j
				break
			}
		}
	}
	fmt.Printf("[%d,%d]", res_i, res_j)

}
