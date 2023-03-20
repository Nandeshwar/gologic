package main

import (
	"fmt"
)

/*
  Move: left to right
   and top to bottom
	{1, 2, 3},
	{1, 1, 6},
	{7, 8, 9},

	min path = 1 + 1+ 1 + 6 + 9 = 18

	1. traverse by row then by column
	2. if somewhere in middle, take minimum of left and top and keep adding

*/
func main() {
	a := [][]int{
		{1, 2, 3},
		{1, 1, 6},
		{7, 8, 9},
	}

	b := [][]int{
		{1, 2, 3},
		{1, 1, 6},
		{7, 8, 9},
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i-1 >= 0 && j-1 >= 0 {
				a[i][j] += min(a[i-1][j], a[i][j-1])
			} else if i-1 >= 0 {
				a[i][j] += a[i-1][j]
			} else if j-1 >= 0 {
				a[i][j] += a[i][j-1]
			}
		}
	}

	fmt.Println("minimum path=", a[2][2])

	m := 3
	n := 3
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	fmt.Println(minPathSumInGridRec(b, dp, m-1, n-1))
}

func minPathSumInGridRec(a, dp [][]int, m, n int) int {
	if m == 0 && n == 0 {
		return a[m][n]
	}

	if m < 0 || n < 0 {
		return int(^uint(0) >> 2)
	}

	if dp[m][n] != -1 {
		return dp[m][n]
	}

	top := a[m][n] + minPathSumInGridRec(a, dp, m-1, n)
	left := a[m][n] + minPathSumInGridRec(a, dp, m, n-1)

	dp[m][n] = min(top, left)
	return min(top, left)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
