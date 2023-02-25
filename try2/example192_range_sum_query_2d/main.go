package main

import (
	"fmt"
)

/*
	{1, 2, 3},
	{1, 2, 3},
	{1, 2, 3},

	find sum total sum : (1,1) - (2,2)
	step1: let's create 1 extra row and 1 extra column
	   0 0 0 0
	   0 0 0 0
	   0 0 0 0
	   0 0 0 0

	step2:
	    calculate presum: 1st row is easy
	   0 0 0 0
	   0	 1 3 6
		2nd row
		 current array + leftSum(dp) + top value(dp) - top left daigonal(dp)
	   0 2 6 12
	   0	 3 9 18
	
   step3: 
          
        	 	0 0    0 0
 			0 [1]  3  6 <- 
 			0 2    {6 12
 			0 3     9 18}
			  ^
			18 - 6 - 3 + 1 = 10

*/

func main() {

	a := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	m := len(a) + 1
	n := len(a[0]) + 1

	row1 := 1
	col1 := 1
	row2 := 2
	col2 := 2
	// output: 10

	/*
		a := [][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		}
		m := len(a) + 1
		n := len(a[0]) + 1

		row1 := 2
		col1 := 1
		row2 := 4
		col2 := 3
		// output: 8

		row1 = 1
		col1 = 1
		row2 = 2
		col2 = 2
		// output: 11

		row1 = 1
		col1 = 2
		row2 = 2
		col2 = 4
		// output: 12
	*/

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// curr arr  value + left value +  top value - top diagonal value
			dp[i][j] = a[i-1][j-1] + dp[i][j-1] + dp[i-1][j] - dp[i-1][j-1]
		}
	}
	display(dp)

	fmt.Println(sumRange(dp, row1, col1, row2, col2))
}

func sumRange(dp [][]int, row1, col1, row2, col2 int) int {
	row1++
	col1++
	row2++
	col2++

	return dp[row2][col2] - dp[row1-1][col2] - dp[row2][col1-1] + dp[row1-1][col1-1]
}

func display(a [][]int) {
	fmt.Println()
	for i := 0; i < len(a); i++ {
		fmt.Println()
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(" ", a[i][j])
		}
	}
	fmt.Println()
}
