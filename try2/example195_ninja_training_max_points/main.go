package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		{10, 20, 30},
		{3, 2, 1},
		{1, 5, 2},
		{1, 3, 2},
	}

	/*
		1st day max activity point = 30
		2nd day max activity point = 3
		3rd day max activiy point = 5
		4th day max activiy point = 2( can not pick 3 because privious day picked same activiy
		                               so next day will be different activity

		ouput should be 30 + 3 + 5 + 2 = 40

	*/

	dp := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		dp[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			dp[i][j] = -1
		}
	}
	mp := maxPoints(a, len(a)-1, 3, dp) // 3 means no activity
	fmt.Println("max points=", mp)

	mp = maxPoints2(a)
	fmt.Println(mp)
}

func maxPoints(a [][]int, day, lastActiviy int, dp [][]int) int {
	if day == 0 {
		var m int

		// 3 for activity: there are 3 activities
		for i := 0; i < 3; i++ {
			if i != lastActiviy {
				m = max(m, a[0][i])
			}
		}
		return m
	}

	if dp[day][lastActiviy] != -1 {
		fmt.Println("dp working")
		return dp[day][lastActiviy]
	}

	var m int
	for i := 0; i < 3; i++ {
		if i != lastActiviy {
			points := a[day][i] + maxPoints(a, day-1, i, dp)
			m = max(m, points)
		}
	}
	dp[day][lastActiviy] = m
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPoints2(aa [][]int) int {
	result := 0

	previousI := -1
	for _, a := range aa {
		currentMax := 0
		currentMaxI := 0
		for i, v := range a {
			if i == previousI {
				continue
			}

			if v > currentMax {
				currentMax = v
				currentMaxI = i
			}
		}
		result += currentMax
		previousI = currentMaxI
	}
	return result
}
