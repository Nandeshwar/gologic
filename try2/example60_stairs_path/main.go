package main

import (
	"fmt"
)

func main() {
	allPaths := stairsAllPath(4) // output: [1111 112 121 13 211 22 31]
	fmt.Println(allPaths)
	fmt.Println(stairPath2(4))
}

// can jump 1 or 2 or 3
func stairsAllPath(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n < 0 {
		return []string{}
	}

	paths1 := stairsAllPath(n - 1)
	paths2 := stairsAllPath(n - 2)
	paths3 := stairsAllPath(n - 3)

	var paths []string

	for _, v := range paths1 {
		paths = append(paths, "1"+v)
	}

	for _, v := range paths2 {
		paths = append(paths, "2"+v)
	}

	for _, v := range paths3 {
		paths = append(paths, "3"+v)
	}
	return paths
}

func stairPath2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = dp[0] + dp[1]

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}
