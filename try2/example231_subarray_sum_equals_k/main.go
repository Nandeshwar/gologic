package main

import (
	"fmt"
)

func main() {
	a := []int{1, 1, 1}
	k := 2
	// ----> output 2

	a = []int{1, 2, 3}
	k = 3
	// -----> output 2
	
	
	cnt := subArrayCount(a, k)
	fmt.Println(cnt)
}

func subArrayCount(a []int, k int) int {
	cnt := 0
	currSum := 0
	m := map[int]int{0: 1}

	for _, v := range a {
		currSum += v
		
		// how many currSum-k : if exists will be added to cnt
		// if sum-k exists and chopped, then sub array exists
		currSum_k := m[currSum-k]
		cnt += currSum_k

		m[currSum] = 1 + m[currSum]
	}
	return cnt
}
