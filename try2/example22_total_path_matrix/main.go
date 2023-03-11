package main

import (
	"fmt"
)

func main() {

	m := 5
	n := 5

	// output: 2

	fmt.Println("total path=", countPath(m, n))

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	fmt.Println("total path algo2=", countPath2(m-1, n-1, dp))
	fmt.Println("total path tabulation=", countPathTabulation(m, n))

	fmt.Println("total path algo3=", countPath3(m, n))
}

func countPath2(r, c int, dp [][]int) int {
	if r == 0 && c == 0 {
		return 1
	}

	if r < 0 || c < 0 {
		return 0
	}

	if dp[r][c] != -1 {
		fmt.Println("dp is working")
		return dp[r][c]
	}

	up := countPath2(r-1, c, dp)
	left := countPath2(r, c-1, dp)

	dp[r][c] = up + left
	return up + left
}

func countPathTabulation(r, c int) int {
	dp := make([][]int, r)
	for i := 0; i < r; i++ {
		dp[i] = make([]int, c)
	}

	// if r == 0 && c == 0 {
	// 	return 1
	// }

	// if r < 0 || c < 0 {
	// 	return 0
	// }

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else {
				up := 0
				left := 0

				if i > 0 {
					up = dp[i-1][j]
				}
				if j > 0 {
					left = dp[i][j-1]
				}

				dp[i][j] = up + left
			}
		}

	}

	// up := countPath2(r-1, c, dp)
	// left := countPath2(r, c-1, dp)

	return dp[r-1][c-1]
}

func countPath3(r, c int) int {
	if r == 1 && c == 1 {
		return 1
	}

	if r < 0 || c < 0 {
		return 0
	}

	return countPath(r-1, c) + countPath(r, c-1)
}

func countPath(r, c int) int {
	if r == 1 || c == 1 {
		return 1
	}

	return countPath(r-1, c) + countPath(r, c-1)
}
