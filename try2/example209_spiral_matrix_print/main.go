package main

import (
	"fmt"
)

func display(m [][]int) {
	fmt.Println()
	for i := 0; i < len(m); i++ {
		fmt.Println()
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(" ", m[i][j])
		}
	}
	fmt.Println()
}

func main() {
	m := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}

	display(m)
	printSpiral(m)
	/*

		 1 2 3
		 8 9 4
		 7 6 5
		1
		8
		7
		6
		5
		4
		3
		2
		9
	*/
}

func printSpiral(m [][]int) {
	rMin := 0
	rMax := len(m) - 1
	cMin := 0
	cMax := len(m[0]) - 1

	cntMax := len(m) * len(m[0])
	cntMin := 0
	for cntMin < cntMax {

		//traverse left wall(top->bottom): traverse each row: col fixed to cMin
		for i := rMin; i <= rMax && cntMin < cntMax; i++ {
			fmt.Println(m[i][cMin])
			cntMin++
		}
		// move from left to right column so that we dont print left bottom corner again
		cMin++

		// traverse bottom wall: left >right): traverse each col: row fixed to cMax
		for j := cMin; j <= cMax && cntMin < cntMax; j++ {
			fmt.Println(m[rMax][j])
			cntMin++
		}

		// move from bottom to up row to avoid printing right bottom corner again
		rMax--

		// traverse right wall: move from bottom to top:
		for i := rMax; i >= rMin && cntMin < cntMax; i-- {
			fmt.Println(m[i][cMax])
			cntMin++
		}

		// move from rightTop to leftTop to avoid printing corner once again
		cMax--

		// traverse top wall : move from right to left
		for i := cMax; i >= cMin && cntMin < cntMax; i-- {
			fmt.Println(m[rMin][i])
			cntMin++
		}
		// move from top to bottom row
		rMin++
	}
}
