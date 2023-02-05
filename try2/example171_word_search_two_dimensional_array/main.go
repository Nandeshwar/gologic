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

	word := "ram"

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i][j] == rune(word[0]) {
				if search(a, word, i, j, 0) {

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

func search(a [][]rune, word string, row, col int, count int) bool {

	if count == len(word) {
		return true
	}

	if row >= 3 || col >= 3 || row < 0 || col < 0 || a[row][col] != rune(word[count]) {
		return false
	}

	//tmp := a[row][col]
	return search(a, word, row+1, col, count+1) ||
		search(a, word, row-1, col, count+1) ||
		search(a, word, row, col+1, count+1) ||
		search(a, word, row, col-1, count+1)
	//a[row][col] = tmp

}
