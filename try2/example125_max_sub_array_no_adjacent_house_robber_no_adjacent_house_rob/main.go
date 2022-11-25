package main

import (
	"fmt"
)

func main() {
	//a := []int{1, 2, 3, 1}
	a := []int{1, 2, 4, 1}
	dp := make([]int, len(a))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	lastInd := len(a) - 1

	fmt.Println(maxHouseRob(a, dp, lastInd))
}

func maxHouseRob(a []int, dp []int, lastInd int) int {
	if lastInd == 0 {
		return a[lastInd]
	}

	if lastInd < 0 {
		return 0
	}

	if dp[lastInd] != -1 {
		return dp[lastInd]
	}

	pick := a[lastInd] + maxHouseRob(a, dp, lastInd-2)
	notPick := 0 + maxHouseRob(a, dp, lastInd-1)

	m := max(pick, notPick)
	dp[lastInd] = m
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
output: 5
*/
