/*
1 0 1 
0 1 1


1 is island
0 is water

Max Area is 3 here: as 1 has 2 others links - horizontal and vertical
*/


package main

import (
	"fmt"
	"math"
)

var seen = [2][3]bool{} 

func main() {
	grid := [][]int{
		{1, 0, 1},
		{0, 1, 1},
	}

	max := 0
	

	rowLen := len(grid)
	colLen := len(grid[0])
	
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			max = int(math.Max(float64(max), float64(islandLen(i, j, grid))))
		}
	}

	fmt.Println(max)
}

func islandLen(i, j int, grid [][]int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 ||   seen[i][j] {
		return 0
	}

	seen[i][j] = true
	return 1 + islandLen(i-1, j, grid) + islandLen(i+1, j, grid) + islandLen(i, j-1, grid) + islandLen(i, j+1, grid)
}

// output: 
// 3