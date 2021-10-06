package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	output := 49

	// use two pointers left and right
	// find width right - left
	// find height take two values(left and right) from array and take smaller one
	// find the area 
	// check if this is max

	i := 0;
	j := len(input) - 1

	maxArea := 0
	for i < j {
		var h = 0
		var w = 0

		if input[i] < input[j] {
			h = input[i]
			w = j -i
			i++
		} else {
			h = input[j]
			w = j -i 
			j--
		}
		area := h * w
		maxArea = int(math.Max(float64(maxArea), float64(area)))
	}

	fmt.Println(maxArea)
	fmt.Println(maxArea == output)
}