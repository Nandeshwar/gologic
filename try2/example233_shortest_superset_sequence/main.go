/*
Given two strings str1 and str2, return the shortest string that has both str1 and str2 as subsequences. If there are multiple valid strings, return any of them.

A string s is a subsequence of string t if deleting some number of characters from t (possibly 0) results in the string s.



Example 1:

Input: str1 = "abac", str2 = "cab"
Output: "cabac"
Explanation:
str1 = "abac" is a subsequence of "cabac" because we can delete the first "c".
str2 = "cab" is a subsequence of "cabac" because we can delete the last "ac".
The answer provided is the shortest such string that satisfies these properties.
*/

package main

import (
	"fmt"
)

func main() {
	/*
	   Longest common sequence:
	     c a b  j
	   0 0 0 0 
i	 a 0 0 1 1 
	 b 0 0 0 2 
	 a 0 0 1 2 
	 c 0 1 1 2

	*/
	str1 := "abac"
	str2 := "cab"

	// output: cabac

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

	fmt.Println("longest common subsequence count=", a[m][n])
	fmt.Println("Finding shortest superset")

	/*
	 Longest common sequence:
	     c a b      <- j
	   0 0 0 0 
i	 a 0 0 1 1 
	 b 0 0 0 2 
	 a 0 0 1 2 
	 c 0 1 1 2
	
	// output: cabac
	*/

	i := m
	j := n
	var result []byte
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			result = append(result, str1[i-1])
			i--
			j--
		} else if a[i-1][j] > a[i][j-1] {
			result = append(result, str1[i-1])
			i--
		} else if a[i][j-1] > a[i-1][j] {
			result = append(result, str2[j-1])
			j--
		}
	}

	for i > 0 {
		result = append(result, str1[i-1])
		i--
	}

	for j > 0 {
		result = append(result, str2[j-1])
		j--
	}

	// reverse string
	x := 0
	y := len(result) - 1

	for x < y {
		result[x], result[y] = result[y], result[x]
		x++
		y--
	}
	fmt.Println("Shortest super sequence=", string(result))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
