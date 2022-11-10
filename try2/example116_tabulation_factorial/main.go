package main

import (
	"fmt"
)

func main() {
	n := 5
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			dp[i] = i
			continue
		}

		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp[n-1])

	fmt.Println("------")
	for i := 0; i < n; i++ {
		fmt.Println(dp[i])
	}
}
