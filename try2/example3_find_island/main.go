package main

import (
	"fmt"
)

func main() {

	arr := [][]int{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 1},
		{1, 0, 0, 0},
	}

	arr2 := [][]int{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 1},
		{1, 0, 0, 0},
	}

	count := 0

	rowLen := len(arr)
	colLen := len(arr[0])

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if arr[i][j] == 1 {
				count++
				turn1to0(arr, rowLen, colLen, i, j)
			}
		}
	}
	fmt.Println("number of islands=", count)

	maxIsland := 0
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if arr2[i][j] == 1 {
				maxIsland = max(maxIsland, MaxIsland(arr2, rowLen, colLen, i, j))
			}
		}
	}
	fmt.Println("max islands=", maxIsland)
}

func turn1to0(arr [][]int, rowLen, colLen, i, j int) {
	if i < 0 || i >= rowLen || j < 0 || j >= colLen || arr[i][j] == 0 {
		return
	}
	arr[i][j] = 0

	turn1to0(arr, rowLen, colLen, i+1, j)
	turn1to0(arr, rowLen, colLen, i-1, j)
	turn1to0(arr, rowLen, colLen, i, j+1)
	turn1to0(arr, rowLen, colLen, i, j-1)

}

func MaxIsland(arr [][]int, rowLen, colLen, i, j int) int {
	if i < 0 || i >= rowLen || j < 0 || j >= colLen || arr[i][j] == 0 {
		return 0
	}
	arr[i][j] = 0
	count := 1
	count += MaxIsland(arr, rowLen, colLen, i+1, j)
	count += MaxIsland(arr, rowLen, colLen, i-1, j)
	count += MaxIsland(arr, rowLen, colLen, i, j+1)
	count += MaxIsland(arr, rowLen, colLen, i, j-1)
	return count

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
