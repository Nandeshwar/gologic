package main

import (
	"fmt"
)

func main() {
	str1 := "acdefghij"
	str2 := "abcde"

	// output: 4

	m := len(str1)
	n := len(str2)

	a := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		a[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// two logic
			// 1. if letter are common in both string, take top diagonal value + 1
			//   else
			// 2. max of top and left
			if str1[i-1] == str2[j-1] {
				a[i][j] = a[i-1][j-1] + 1
			} else {
				a[i][j] = max(a[i-1][j], a[i][j-1])
			}
		}
	}

	fmt.Println(a)
	fmt.Println(a[m][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
