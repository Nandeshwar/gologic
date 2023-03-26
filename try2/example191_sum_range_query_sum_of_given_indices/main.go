package main

import (
	"fmt"
)

func main() {
	a := []int{-1, -2, -3, 4, 5, 6, -2}
	// dp= [-1 -3 -6 -2 3 9 7]
	// sum(0, 3) = -1 + (-2) + (-3) + 4 = -2
	// sum(2, 4) = -3 + 4 + 5 = 6

	dp := make([]int, len(a))
	dp[0] = a[0]
	for i := 1; i < len(a); i++ {
		dp[i] = dp[i-1] + a[i]
	}

	fmt.Println("dp=", dp)
	fmt.Println(sumRange(0, 3, dp))
	fmt.Println(sumRange(2, 4, dp))
}

func sumRange(i, j int, dp []int) int {

	if i == 0 {
		return dp[j]
	}

	return dp[j] - dp[i-1]
}
