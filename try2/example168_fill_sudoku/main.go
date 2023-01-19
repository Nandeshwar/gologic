package main

import (
	"fmt"
)

type Ints [9][9]int

func (a Ints) display() {
	fmt.Println()
	for i := 0; i < 9; i++ {
		fmt.Println()
		for j := 0; j < 9; j++ {
			fmt.Print(" ", a[i][j])
		}
	}
	fmt.Println()
}

/*
output:
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 1 0 0 0 0 0
 0 0 0 0 0 2 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0


 1 2 3 4 5 6 7 8 9
 4 5 6 7 8 9 1 2 3
 7 8 9 2 1 3 4 5 6
 2 3 4 1 6 5 8 9 7
 5 1 7 8 9 2 3 6 4
 6 9 8 3 4 7 2 1 5
 3 4 5 6 2 1 9 7 8
 8 6 1 9 7 4 5 3 2
 9 7 2 5 3 8 6 4 1
*/

func main() {
	var board Ints
	board[3][3] = 1
	board[4][5] = 2

	board.display()

	fillSudoku(&board)
	board.display()

}

func fillSudoku(a *Ints) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*a)[i][j] == 0 {
				for k := 1; k <= 9; k++ {
					if canBePlacedInBoard(*a, i, j, k) {
						(*a)[i][j] = k
						if fillSudoku(a) == true {
							return true
						} else {
							(*a)[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func canBePlacedInBoard(a Ints, i, j, k int) bool {
	for x := 0; x < 9; x++ {
		// check entire row
		if a[i][x] == k {
			return false
		}

		// check entire column
		if a[x][j] == k {
			return false
		}

		// check in 3*3 box
		if a[3*(i/3)+x/3][3*(j/3)+x%3] == k {
			return false
		}
	}
	return true
}
