package main

import (
	"fmt"
)

/*
	1st and last house can not be robbed
	prepare two arrays: one has everything but last
	other array has everything but first

	and find max value robbed from 1st array
	and find max value robbed from 2nd array

	At end find max of two max
*/

func main() {
	a := []int{20, 1, 3, 4} // expected answer should be  23
	a1 := a[1:]
	a2 := a[:len(a)-1]

	dp1 := make([]int, len(a1))
	dp2 := make([]int, len(a2))

	for i := 0; i < len(dp1); i++ {
		dp1[i] = -1
		dp2[i] = -1
	}
	m1 := houseRob(a1, len(a1)-1, dp1)
	m2 := houseRob(a2, len(a2)-1, dp2)

	fmt.Println(max(m1, m2))
}

func houseRob(a []int, lastInd int, dp []int) int {
	if lastInd == 0 {
		return a[0]
	}
	if lastInd < 0 {
		return 0
	}
	if dp[lastInd] != -1 {
		return a[lastInd]
	}

	pick := a[lastInd] + houseRob(a, lastInd-2, dp)
	notPick := 0 + houseRob(a, lastInd-1, dp)

	m := max(pick, notPick)
	dp[lastInd] = m
	return m

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
