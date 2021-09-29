package main

import (
	"fmt"
)

const EMPTY = "_"
func main() {
	board := [3][3]string{
		{EMPTY,EMPTY,EMPTY},
		{EMPTY,EMPTY,EMPTY},
		{EMPTY,EMPTY,EMPTY},
	}

	for {
		x := 0
		y := 0
		if boardFilled(&board) {
			fmt.Println("no one is winner")
			break
		}
		displayBoard(&board)
		fmt.Println("\n player1:")
		fmt.Scanf("%d %d", &x, &y)
		playX(x, y, &board)
		if gameOver(&board) {
			fmt.Println("Player1 is winner")
			break
		}

		displayBoard(&board)
		fmt.Println("\n player2:")
		fmt.Scanf("%d %d", &x, &y)
		playO(x, y, &board)
		if gameOver(&board) {
			fmt.Println("Player2 is winner")
			break
		}
	}
	


}

func checkNextMoveAvailability(x, y int, board *[3][3]string) bool {
	if board[x][y] == EMPTY {
		return true
	}
	return false
}

func playX(x, y int, board *[3][3]string) bool {
	fmt.Println("Player1")
	if !checkNextMoveAvailability(x, y, board) {
		fmt.Println("already filled ")
		return false
	} 

	board[x][y] = "X"
	return true
}

func playO(x, y int, board *[3][3]string) bool {
	fmt.Println("Player2")
	if !checkNextMoveAvailability(x, y, board) {
		fmt.Println("already filled ")
		return false
	} 

	board[x][y] = "O"
	return true
}

func displayBoard(board *[3][3]string) {
	for i := 0; i < 3; i++ {
		fmt.Println()
		for j := 0; j < 3; j++ {
			fmt.Printf("%s ", board[i][j]);
		}
	}
}

func gameOver(board *[3][3]string) bool {
	// check row is filled
	for i := 0; i < 3; i++ {
		if (board[i][0] == "X" && board[i][1] == "X" && board[i][2] == "X") || (board[i][0] == "O" && board[i][1] == "O" && board[i][2] == "O") {
			return true
		}
	}

	// check col is filled
	for i := 0; i < 3; i++ {
		if (board[0][i] == "X" && board[1][i] == "X" && board[2][i] == "X") || (board[0][i] == "O" && board[1][i] == "O" && board[2][i] == "O") {
			return true
		}
	}

	// check diaogonal filled
	if board[0][0] == "X" && board[1][1] == "X" && board[2][2] == "X" || board[0][0] == "O" && board[1][1] == "O" && board[2][2] == "O" {
		return true
	}
	if board[2][0] == "X" && board[1][1] == "X" && board[2][2] == "X" || board[2][0] == "O" && board[1][1] == "O" && board[2][2] == "O" {
		return true
	}


	return false;
}

func boardFilled(board *[3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "_" {
				return false;
			}
		}
	}
	return true;
}