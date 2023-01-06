package main

import (
	"fmt"
	"strconv"
)

func main() {
	//str := "21123" // expectation 8
	str := "110"
	count := decodingWaysCount(str) //
	fmt.Println("count=", count)
	/*
		output:
		  11
		[1 2 1]
		count= 1
		bash-3.2$
	*/
}

func decodingWaysCount(a string) int {
	dp := make([]int, len(a))
	dp[0] = 1

	for i := 1; i < len(a); i++ {
		if a[i-1] == '0' && a[i] == '0' { // 100
			dp[i] = 0
		} else if a[i-1] == '0' && a[i] != '0' { // ex 10
			dp[i] = dp[i-1]
		} else if a[i-1] != '0' && a[i] == '0' { // 110
			if a[i-1] == '1' || a[i-1] == '2' {
				if i >= 2 {
					dp[i] = dp[i-2]
				} else {
					dp[i] = 1
				}
			} else {
				dp[i] = 0
			}
		} else {
			twoDigitStr := string(a[i-1]) + string(a[i])
			twoDigit, err := strconv.Atoi(twoDigitStr)
			if err != nil {
				fmt.Println(err)
				return -1
			}

			if twoDigit <= 26 {
				fmt.Println(twoDigit)
				if i >= 2 {
					dp[i] = dp[i-1] + dp[i-2]
				} else {
					dp[i] = dp[i-1] + 1
				}
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[len(a)-1]
}
