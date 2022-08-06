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
