package main

import (
	"fmt"

	"github.com/logic-building/functional-go/fp"
)

// number of coins required to make given sum
func coinCountToMakeSum(coins []int, sum int) int {
	m := map[int]int{}

	coinCount := coinCount(coins, sum, m)
	fmt.Println("m: ", m)
	return coinCount
}

func coinCount(coins []int, sum int, m map[int]int) int {
	var ans = 10000

	if sum == 0 {
		return 0
	}

	for _, c := range coins {
		if sum-c >= 0 {
			v, ok := m[sum]
			if ok {
				return v
			}
			subAns := coinCount(coins, sum-c, m) + 1
			ans = fp.MinInt([]int{subAns, ans})
		}
	}
	m[sum] = ans
	return ans
}
