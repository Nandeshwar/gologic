package main

import (
	"fmt"
)

func main() {

	a := []int{0, 1, 1, 0, 0, 0, 0, 0, 0, 1}
	// output: 4
	
	fmt.Println(equalCount_0_and_1(a))
}

// Algorithm:
// treat 0 as -1
// treat 1 as +1

// store sum as key, and index as value
// if index exist in map, then i - ind  will go to answer(max will be returned at end)
func equalCount_0_and_1(a []int) int {
	m := map[int]int{}
	sum := 0

	maxCnt := 0

	m[0] = -1
	for i := 0; i < len(a); i++ {
		switch a[i] {
		case 0:
			sum += -1
		case 1:
			sum += 1
		}

		ind, ok := m[sum]
		if ok {
			cnt := i - ind
			maxCnt = max(maxCnt, cnt)
		} else {
			m[sum] = i
		}
	}
	return maxCnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
