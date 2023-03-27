package main

import (
	"fmt"
)

func main() {
	a := []int{1, 4, 4, 4} // output: 2, 3

	result := findAllMissingNumbersAlgo1(a)
	fmt.Println(result)

	// Algorithm
	// 1. iterate array and find index of item which is item-1
	// 2. update new index with negative of item
	// 3. iterate array again
	// 4. for positive value: find index +1 will be added to result
	result = findAllMissingNumbersAlgo2(a)
	fmt.Println(result)
}

func findAllMissingNumbersAlgo1(a []int) []int {
	var result []int

	m := make(map[int]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}

	for i := 1; i <= len(a); i++ {
		_, ok := m[i]
		if !ok {
			result = append(result, i)
		}
	}
	return result
}

func findAllMissingNumbersAlgo2(a []int) []int {
	var result []int
	for _, v := range a {
		if v > 0 {
			indOfItem := abs(v) - 1
			a[indOfItem] = -1 * abs(v)
		}
	}

	for ind, v := range a {
		if v >= 0 {
			result = append(result, ind+1)
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
