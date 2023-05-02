package main

import (
	"fmt"
)

func main() {
	/*
		observation:
		------------
		every row is in ascending order with 0, 1

		Algorithm:
		----------
		1. Travese from bottom
		2. if 1 move to left and update result
		3. if 0 move to top and update result


	*/
	m := [][]int{
		{0, 0, 0, 1},
		{0, 0, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 1},
	}

	fmt.Println(findLeftMostColumnWithAtLeastOne1(m)) // output: 2 (3rd column)
}

func findLeftMostColumnWithAtLeastOne1(m [][]int) int {
	result := -1

	i := len(m) - 1
	j := len(m[i]) - 1

	for {
		if i < 0 || j < 0 {
			break
		}

		if m[i][j] == 1 {
			result = j
			j-- // move left
		} else if m[i][j] == 0 {
			i-- // move up
		}
	}
	return result

}
