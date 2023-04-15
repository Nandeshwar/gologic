package main

import (
	"fmt"
)

func main() {
	// Find shortest unsorted contiguous subarray : input is in ascending order
	a := []int{2, 6, 4, 8, 10, 9, 15}
	// output: 6, 4, 8, 10, 9

	fmt.Println(findShortestUnsortedSubArray(a))
}

func findShortestUnsortedSubArray(a []int) []int {

	startIndex := -1
	for i := 0; i < len(a)-1; i++ {
		if a[i+1] < a[i] {
			startIndex = i
			break
		}
	}

	endIndex := -1
	for i := len(a) - 1; i > 0; i-- {
		if a[i-1] > a[i] {
			endIndex = i
			break
		}
	}

	if startIndex == -1 || endIndex == -1 {
		return []int{}
	}
	return a[startIndex : endIndex+1]
}
