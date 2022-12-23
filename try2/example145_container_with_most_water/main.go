package main

import (
	"fmt"
)

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("max area of container =", bruteForce(a))
	fmt.Println("max area of container =", maxAreaContainer(a))
	/*
		output:
		max area of container = 49
		max area of container = 49
	*/
}

func bruteForce(a []int) int {
	var maxArea int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			// width * height
			// height: min of left and right boundary
			area := (j - i) * min(a[i], a[j])
			maxArea = max(maxArea, area)
		}
	}
	return maxArea
}

func maxAreaContainer(a []int) int {
	maxArea := 0
	i := 0
	j := len(a) - 1

	for i < j {
		area := (j - i) * min(a[i], a[j])
		maxArea = max(maxArea, area)

		if a[i] < a[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
