package main

import (
	"testing"
	"fmt"
)
var board [3][3]string = [3][3]string{
	{EMPTY,EMPTY,EMPTY},
	{EMPTY,EMPTY,EMPTY},
	{EMPTY,EMPTY,EMPTY},
}

func TestWinner(t *testing.T) {
	board[0][0] = "X"
	board[1][1] = "X"
	board[2][2] = "X"

	expectedResult := true
	result := gameOver(&board)
	if result != expectedResult {
		t.Errorf("diagonal value is not set")
	}
}

func TestPlayx(t *testing.T) {
	board[0][0] = "O"
	board[1][1] = "O"
	board[2][2] = "O"

	expectedResult := true
	result := gameOver(&board)
	fmt.Println(result)
	if result != expectedResult {
		t.Errorf("diagonal value is not set")
	}
}

