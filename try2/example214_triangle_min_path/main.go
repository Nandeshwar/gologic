package main

import (
	"fmt"
)

func main() {
	// move bottom and diagonal
	triangle := [][]int{
		{1},
		{2, 3},
		{3, 6, 7},
		{8, 9, 6, 10},
	}

	dp := make([][]int, len(triangle))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = -1
		}
	}
	fmt.Println(minPath(triangle, 0, 0, dp))
}

func minPath(aa [][]int, i, j int, dp [][]int) int {
	m := len(aa)

	if i == m-1 {
		return aa[i][j]
	}

	if dp[i][j] != -1 {
		fmt.Println("dp working")
		return dp[i][j]
	}

	sum1 := aa[i][j] + minPath(aa, i+1, j, dp)
	sum2 := aa[i][j] + minPath(aa, i+1, j+1, dp)

	dp[i][j] = min(sum1, sum2)
	return min(sum1, sum2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
