package main

import (
	"fmt"
)

type Matrix [][]int

func (m *Matrix) display() {
	r := len(*m)
	fmt.Println("")
	for i := 0; i < r; i++ {
		c := len((*m)[i])
		fmt.Println("")
		for j := 0; j < c; j++ {
			fmt.Print(" ", (*m)[i][j])
		}
		fmt.Println()
	}
}

func main() {
	a := Matrix{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	// output: reverse each row and flip bit
	/*
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
		  |
		  |
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},

	*/

	a.display()
	flipImage(&a)
	a.display()

}

func flipImage(a *Matrix) {
	r := len(*a)

	for i := 0; i < r; i++ {
		j := 0
		k := len((*a)[i]) - 1

		for j < k {
			(*a)[i][j], (*a)[i][k] = (*a)[i][k], (*a)[i][j]
			j++
			k--
		}

		for j = 0; j < len(*a); j++ {
			if (*a)[i][j] == 1 {
				(*a)[i][j] = 0
			} else {
				(*a)[i][j] = 1
			}
		}
	}
}
