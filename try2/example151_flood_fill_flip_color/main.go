package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		{1, 1, 1},
		{1, 1, 3},
		{2, 3, 1},
	}

	// if old color 1 is selected and new color 5 is selected,
	// all adjusent 1 will be replaces with 5
	// output
	/*
		{5, 5, 5},
		{5, 5, 3},
		{2, 3, 1},
	*/

	oldColor := 1
	newColor := 5
	initLocationX := 0
	initLocationY := 0

	floodFill(&a, initLocationX, initLocationY, oldColor, newColor, 3, 3)

	fmt.Println("")
	for i := 0; i < 3; i++ {
		fmt.Println("")
		for j := 0; j < 3; j++ {
			fmt.Print(" ", a[i][j])
		}
	}
	fmt.Println()

	b := [][]int{
		{1, 1, 1},
		{1, 1, 3},
		{2, 3, 1},
	}

	deltaRow := []int{-1, 1, 0, 0}
	deltaCol := []int{0, 0, -1, 1}

	// deltaRow := []int{-1, 0, +1, 0}
	// deltaCol := []int{0, +1, 0, -1}

	b[initLocationX][initLocationY] = newColor

	floodFill2(&b, initLocationX, initLocationY, oldColor, newColor, 3, 3, deltaRow, deltaCol)

	fmt.Println("")
	for i := 0; i < 3; i++ {
		fmt.Println("")
		for j := 0; j < 3; j++ {
			fmt.Print(" ", b[i][j])
		}
	}
	fmt.Println()
}

func floodFill(a *[][]int, i, j, oldColor, newColor, rl, cl int) {
	if i < 0 || j < 0 || i >= rl || j >= cl || (*a)[i][j] != oldColor || (*a)[i][j] == newColor {
		return
	}

	if (*a)[i][j] == oldColor {
		(*a)[i][j] = newColor
	}

	floodFill(a, i+1, j, oldColor, newColor, 3, 3)
	floodFill(a, i-1, j, oldColor, newColor, 3, 3)
	floodFill(a, i, j+1, oldColor, newColor, 3, 3)
	floodFill(a, i, j-1, oldColor, newColor, 3, 3)

}

func floodFill2(a *[][]int, i, j, oldColor, newColor, rl, cl int, deltaRow, deltaCol []int) {
	for k := 0; k < 4; k++ {

		if i < 0 || j < 0 || i >= rl || j >= cl || (*a)[i][j] != oldColor || (*a)[i][j] == newColor {
			return
		}

		if (*a)[i][j] == oldColor {
			(*a)[i][j] = newColor
		}
		floodFill2(a, i+deltaRow[k], j+deltaCol[k], oldColor, newColor, 3, 3, deltaRow, deltaCol)
	}
}
