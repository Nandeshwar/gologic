package main

import (
	"fmt"
)

const n = 4

func main() {
	board := [n][n]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = 0
		}
	}
	fmt.Println("\n", nqueen(board, 0))
}

func nqueen(board [n][n]int, col int) bool {
	if col >= n {
		displayBoard(board)
		return true
	}

	for i := 0; i < n; i++ {

		if isGoodPlaceForQueen(board, i, col) {
			board[i][col] = 1

			if nqueen(board, col+1) == true {
				return true
			}
			board[i][col] = 0
		}
	}
	return false
}

func isGoodPlaceForQueen(board [n][n]int, row, col int) bool {
	// check in left side
	for i := col; i >= 0; i-- {
		if board[row][i] == 1 {
			return false
		}
	}

	// check in upper left diagonal

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// check in lower diagonal
	for i, j := row, col; i < n && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func displayBoard(board [n][n]int) {
	for i := 0; i < n; i++ {
		fmt.Println()
		for j := 0; j < n; j++ {
			fmt.Print("\t", board[i][j])
		}
	}
}

/*
output:
	0	0	1	0
	1	0	0	0
	0	0	0	1
	0	1	0	0
 true
*/