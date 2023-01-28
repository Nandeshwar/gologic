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
	minSteps := frogJump(a, len(a)-1, dp)
	fmt.Println("min steps=", minSteps)
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
