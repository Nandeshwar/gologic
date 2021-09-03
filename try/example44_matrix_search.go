package main

import (
	"fmt"
)
func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8 , 9},
	}
	if findItem(6, matrix) {
		fmt.Println("Item found")
	} else {
		fmt.Println("Item not found")
	}
}

func findItem(num int, matrix [][]int) bool {
	rowLen := len(matrix)
	colLen := len(matrix[0])

	matrixLen := (rowLen * colLen)
	left := 0
	right := matrixLen - 1;

	for left <= right {
		midPoint := left + (right - left) / 2
		midPointElement := matrix[midPoint / colLen][midPoint % colLen]

		if midPointElement == num {
			return true
		} else if (num < midPointElement) {
			right = midPoint -1
		} else {
			left = midPoint + 1
		}
	}
	
	return false
}