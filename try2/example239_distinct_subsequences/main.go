package main

import (
	"fmt"
)

func main() {
	s1 := "abcdabc"
	s2 := "abc"

	// occurence of abc unique count = 4
	/*
		abc dabc
		a bcda  bc
		ab bcda c
		abcd abc
	*/
	i := len(s1) - 1
	j := len(s2) - 1

	dp := make([][]int, len(s1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2))
		for j := 0; j < len(s2); j++ {
			dp[i][j] = -1
		}
	}
	fmt.Println(distinctSubsequence(s1, s2, i, j, dp))
}

func distinctSubsequence(s1, s2 string, i, j int, dp [][]int) int {
	if j < 0 {
		return 1 // j < 0 that means found
	}

	if i < 0 {
		return 0
	}

	if dp[i][j] != -1 {
		fmt.Println("dynamic programming working")
		return dp[i][j]
	}

	if s1[i] == s2[j] {

		cnt := distinctSubsequence(s1, s2, i-1, j-1, dp) + distinctSubsequence(s1, s2, i-1, j, dp) // i-1, j : fix substring and check other character in 1st string
		dp[i][j] = cnt
		return cnt
	}

	cnt := distinctSubsequence(s1, s2, i-1, j, dp) // i-1, j : fix with s2 and try to match with all characters of s1
	dp[i][j] = cnt
	return cnt
}
