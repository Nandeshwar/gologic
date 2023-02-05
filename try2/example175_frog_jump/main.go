package main

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30, 10} // expected output: 20
	/*
		  Frog can jump either 1 step or 2 step.
		  Find miminum energy to reach from beg to end
		  try1:
		    10 20 30 10
			10 -> 20: energy : 10
			20 -> 30: energy: 10
			30 - 10(end): energy: 20
			            Total enerty = 40

			try2:
			10 20 30 10
			10 -> 30: 20
			30->10(end) = 10
			     Total= 30

			try3:
			10 20 30 10
			10-> 20 : 10
			20 -> 10 (end): 10
			           total = 20



		So minimum energy required is = 20
		so output = 20

		Solution approach: Recursion
		Try all possbile way and find minimum

	*/
	dp := make([]int, len(a))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	dp3 := make([]int, len(a))
	for i := 0; i < len(dp3); i++ {
		dp3[i] = -1
	}
	minSteps := frogJump(a, len(a)-1, dp)
	fmt.Println("min steps=", minSteps)

	minSteps2 := frogJump2(a)
	fmt.Println("bottom up approach = min steps=", minSteps2)

	minSteps3 := frogJump3(a, 0, dp3)
	fmt.Println("approach3 = min steps=", minSteps3)
}

func frogJump(a []int, ind int, dp []int) int {
	if ind == 0 {
		return 0
	}

	if dp[ind] > -1 {
		fmt.Println("dp is helping")
		return dp[ind]
	}

	left := frogJump(a, ind-1, dp) + abs(a[ind]-a[ind-1])
	fmt.Println("left=", left)
	right := int(^uint(0) >> 1)
	if ind > 1 {
		right = frogJump(a, ind-2, dp) + +abs(a[ind]-a[ind-2])
	}
	fmt.Println("right=", right)
	m := min(left, right)
	dp[ind] = m
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func frogJump2(a []int) int {
	dp := make([]int, len(a))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	// if ind == 0 {
	// 	return 0
	// }

	dp[0] = 0

	// left := frogJump(a, ind-1, dp) + abs(a[ind]-a[ind-1])
	// right := int(^uint(0) >> 1)
	// if ind > 1 {
	// 	right = frogJump(a, ind-2, dp) + +abs(a[ind]-a[ind-2])
	// }
	// m := min(left, right)

	for i := 1; i < len(dp); i++ {
		left := dp[i-1] + abs(a[i]-a[i-1])
		right := int(^uint(0) >> 1)
		if i > 1 {
			right = dp[i-2] + abs(a[i]-a[i-2])
		}
		dp[i] = min(left, right)
	}

	return dp[len(a)-1]
}

func frogJump3(a []int, ind int, dp []int) int {
	if ind >= len(a) {
		return 0
	}

	if dp[ind] > -1 {
		fmt.Println("dp is helping")
		return dp[ind]
	}

	var left int
	if ind < len(a)-1 {

		left = frogJump3(a, ind+1, dp) + abs(a[ind+1]-a[ind])
	}
	fmt.Println("left=", left)
	right := int(^uint(0) >> 1)
	if ind < len(a)-2 {
		right = frogJump3(a, ind+2, dp) + abs(a[ind+2]-a[ind])
	}
	fmt.Println("right=", right)
	m := min(left, right)
	dp[ind] = m
	return m
}
