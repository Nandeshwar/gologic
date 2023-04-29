/*
You are given a m x n 2D grid initialized with these three possible values.

-1 - A wall or an obstacle. 0 - A gate. INF - Infinity means an empty room. We use the value 231 - 1 = 2147483647 to represent INF as you may assume that the distance to a gate is less than 2147483647. Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, it should be filled with INF.

For example, given the 2D grid:

INF  -1  0  INF
INF INF INF  -1
INF  -1 INF  -1
  0  -1 INF INF
After running your function, the 2D grid should be:

  3  -1   0   1
  2   2   1  -1
  1  -1   2  -1
  0  -1   3   4
*/

package main

import (
	"fmt"
)

// 0 - gate
// -1 - wall
// 1000 - empty space
// find distance from gate(0) to empty space
func main() {
	a := [][]int{
		{1000, -1, 0, 1000},
		{1000, 1000, 1000, -1},
		{1000, -1, 1000, -1},
		{0, -1, 1000, 1000},
	}

	wallAndGates(a)
	for i := 0; i < len(a); i++ {
		fmt.Println()
		for j := 0; j < len(a[i]); j++ {
			fmt.Print("   ", a[i][j])
		}
	}
	fmt.Println()
}

func wallAndGates(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == 0 {
				dfs(a, i, j, 0)
			}
		}
	}
}

func dfs(a [][]int, i, j, counter int) {
	// 1. check out of boundary
	// 2. check if any cell has already value less than counter(i.e already calcuated from some direction)
	if i < 0 || j < 0 || i >= len(a) || j >= len(a[0]) || a[i][j] == -1 || a[i][j] < counter {
		return
	}

	a[i][j] = counter
	dfs(a, i+1, j, counter+1)
	dfs(a, i-1, j, counter+1)
	dfs(a, i, j+1, counter+1)
	dfs(a, i, j-1, counter+1)
}
