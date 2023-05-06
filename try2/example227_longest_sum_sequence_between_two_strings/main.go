package main

import (
	"fmt"
)

func main() {
	str1 := "abcde"
	str2 := "acdefghij"
	
	/*
		Algorithm
		-----------
		1. start with last index of both string
		2. if last letter of both string matches then 1 + function call
		3. if last letter of both string does not match, then
		   max of f(ind1-1, ind2), f(ind1, ind2-1)
	*/

	ind1 := len(str1) - 1
	ind2 := len(str2) - 1

	maxLen := max(len(str1), len(str2))

	dp := make([][]int, maxLen)

	for i := 0; i < maxLen; i++ {
		dp[i] = make([]int, maxLen)
		for j := 0; j < maxLen; j++ {
			dp[i][j] = -1
		}
	}

	fmt.Println(maxSubSequence(str1, str2, ind1, ind2, dp))
}

func maxSubSequence(str1, str2 string, ind1, ind2 int, dp [][]int) int {
	if ind1 == 0 || ind2 == 0 {
		return 0
	}

	if dp[ind1][ind2] != -1 {
		fmt.Println("dp is working...")
		return dp[ind1][ind2]
	}

	if str1[ind1] == str2[ind2] {
		return 1 + maxSubSequence(str1, str2, ind1-1, ind2-1, dp)
	}

	result := max(maxSubSequence(str1, str2, ind1-1, ind2, dp), maxSubSequence(str1, str2, ind1, ind2-1, dp))
	dp[ind1][ind2] = result
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
