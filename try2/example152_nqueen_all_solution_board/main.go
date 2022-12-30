package main

import "fmt"

func main() {
	board := [4][4]byte{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			board[i][j] = '-'
		}
	}

	result := [][4][4]byte{}
	nqueen(&board, 0, &result)
	fmt.Println("result=")
	for i := 0; i < len(result); i++ {
		display(result[i])
	}

}

func nqueen(board *[4][4]byte, col int, result *[][4][4]byte) {
	if col == 4 {
		display(*board)
		
		*result = append(*result, *board)
		return
	}

	for i := 0; i < 4; i++ {
		if rightPlace(*board, i, col) {
			(*board)[i][col] = 'Q'

			nqueen(board, col+1, result)

			(*board)[i][col] = '-'
		}
	}
}

func display(board [4][4]byte) {

	fmt.Println("")
	for i := 0; i < 4; i++ {
		fmt.Println("")
		for j := 0; j < 4; j++ {
			fmt.Print(" ", string(board[i][j]))
		}
	}
	fmt.Println("")
}

func rightPlace(board [4][4]byte, row, col int) bool {
	/*
			  -
			    -
		. .	. . .  -
			    -
		    -
	*/
	// upper left diagonal check
	i := row
	j := col
	for i >= 0 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	// left check
	i = row
	j = col

	for j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		j--
	}

	// lower left diagonal check
	i = row
	j = col

	for i < 4 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i++
		j--
	}
	return true
}

/*
output:
  - - Q -
 Q - - -
 - - - Q
 - Q - -


 - Q - -
 - - - Q
 Q - - -
 - - Q -
*/
