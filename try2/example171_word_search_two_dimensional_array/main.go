package main

import (
	"fmt"
)

func main() {
	a := [][]rune{
		{'a', 'm', 'm'},
		{'k', 'a', 'r'},
		{'m', 'o', 'p'},
	}

	b := [][]byte{
		{'a', 'm', 'm'},
		{'k', 'a', 'r'},
		{'m', 'o', 'p'},
	}

	word := "ram"
out:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i][j] == rune(word[0]) {
				if search(a, word, i, j, 0) {

					fmt.Println("true")
					break out
				} else {
					fmt.Println("false")
					break out
				}
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i][j] == rune(word[0]) {
				if checkInAllDirections(b, word, i, j, 0) {

					fmt.Println("true")
					return
				} else {
					fmt.Println("false")
					return
				}
			}
		}
	}
}
func checkInAllDirections(board [][]byte, word string, i, j int, count int) bool {

	if count == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != word[count] {
		return false
	}

	tmp := board[i][j]
	board[i][j] = ' '
	count1 := checkInAllDirections(board, word, i+1, j, count+1)
	count2 := checkInAllDirections(board, word, i-1, j, count+1)
	count3 := checkInAllDirections(board, word, i, j+1, count+1)
	count4 := checkInAllDirections(board, word, i, j-1, count+1)
	board[i][j] = tmp

	return count1 || count2 || count3 || count4
}

func search(a [][]rune, word string, row, col int, count int) bool {

	if count == len(word) {
		return true
	}

	if row >= 3 || col >= 3 || row < 0 || col < 0 || a[row][col] != rune(word[count]) {
		return false
	}

	tmp := a[row][col]
	a[row][col] = ' '
	found := search(a, word, row+1, col, count+1) ||
		search(a, word, row-1, col, count+1) ||
		search(a, word, row, col+1, count+1) ||
		search(a, word, row, col-1, count+1)
	a[row][col] = tmp
	return found

}
