/*
input: 6, 20, 10, 15, -10, 10
target sum : 5

output: index of 15 and -10
*/
package main

import (
	"fmt"
)

func main() {
	a, b := targetSubArrIndex([]int{6, 20, 10, 15, -10, 10}, 5)
	fmt.Println(a, " ", b)
}

func targetSubArrIndex(a []int, targetSum int) (int, int) {
	currSum := 0

	start := -1
	end := -1
	m := map[int]int{}
	for i := 0; i < len(a); i++ {
		currSum += a[i]
		if currSum-targetSum == 0 {
			start = 0
			end = i
			return start, end
		}

		index1, ok := m[currSum-targetSum]
		if ok {
			start = index1 + 1
			end = i
			return start, end
		}

		m[currSum] = i
	}

	return start, end
}
