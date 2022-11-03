package main

import (
	"fmt"
)

/*
  Move: left to right
   and top to bottom
	{1, 2, 3},
	{1, 1, 6},
	{7, 8, 9},

	min path = 1 + 1+ 1 + 6 + 9 = 18

	1. traverse by row then by column
	2. if somewhere in middle, take minimum of left and top and keep adding

*/
func main() {
	a := [][]int{
		{1, 2, 3},
		{1, 1, 6},
		{7, 8, 9},
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i-1 >= 0 && j-1 >= 0 {
				a[i][j] += min(a[i-1][j], a[i][j-1])
			} else if i-1 >= 0 {
				a[i][j] += a[i-1][j]
			} else if j-1 >= 0 {
				a[i][j] += a[i][j-1]
			}
		}
	}

	fmt.Println("minimum path=", a[2][2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
