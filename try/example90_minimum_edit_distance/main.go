// https://www.youtube.com/watch?v=AuYujVj646Q

package main

import (
	"fmt"
	"math"
)

func main() {

	s1 := "abc"
	s2 := "bcd"
	//s1 = ""

	actCount := min_edit_distance(s1, s2)
	fmt.Println("action count=", actCount)

}

func min_edit_distance(s1 string, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)

	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	s1CharArr := []rune(s1)
	s2CharArr := []rune(s2)

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {

			// if s1 is empty
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if s1CharArr[i-1] == s2CharArr[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {

				d := dp[i-1][j]

				in := dp[i][j-1]

				r := dp[i-1][j-1]

				dp[i][j] = 1 + int(math.Min(math.Min(float64(d), float64(in)), float64(r)))
			}

		}
	}

	for i := 0; i <= l1; i++ {
		fmt.Println()
		for j := 0; j <= l2; j++ {
			fmt.Print(dp[i][j], " ")
		}
	}

	return dp[l1][l2]

}
