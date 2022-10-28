package main

import (
	"fmt"
)

const n = 5

func main() {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	knightMove(a, 0, 0, 1)
}

func knightMove(a [][]int, i, j, move int) {

	if i < 0 || j < 0 || i >= n || j >= n || a[i][j] > 0 {
		return
	}
	if move == n*n {
		a[i][j] = move
		display(a)
		//a[i][j] = 0
		return
	}

	a[i][j] = move
	knightMove(a, i-2, j+1, move+1)
	knightMove(a, i-1, j+2, move+1)
	knightMove(a, i+1, j+2, move+1)
	knightMove(a, i+2, j+1, move+1)
	knightMove(a, i+2, j-1, move+1)
	knightMove(a, i+1, j-2, move+1)
	knightMove(a, i-1, j-2, move+1)
	knightMove(a, i-2, j-1, move+1)
	a[i][j] = 0
}

func display(a [][]int) {
	for i := 0; i < n; i++ {
		fmt.Println()
		for j := 0; j < n; j++ {
			fmt.Print(a[i][j], "\t")
		}
	}
}
