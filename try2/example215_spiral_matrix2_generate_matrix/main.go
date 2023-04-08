package main

import (
	"fmt"
)

func main() {
	m := 3
	n := 3
	/*
		1 2 3
		8 9 4
		7 6	5
	*/
	a := generateSpiralMatrix(m, n)
	display(m, n, a)
}

func generateSpiralMatrix(m, n int) [][]int {

	a := make([][]int, m*n)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
	}

	num := 1
	rMin := 0
	rMax := m - 1

	cMin := 0
	cMax := n - 1

	numMax := m * n

	for num <= numMax {

		// top wall
		for j := 0; j <= cMax && num <= numMax; j++ {
			fmt.Println("top wall=", num)
			a[rMin][j] = num
			num++
		}
		rMin++

		// right wall
		for i := rMin; i <= rMax && num <= numMax; i++ {
			fmt.Println("right wall=", num)
			a[i][cMax] = num
			num++
		}

		cMax--

		// bottom wall
		for j := cMax; j >= cMin && num <= numMax; j-- {
			fmt.Println("bottom wall=", num)
			a[rMax][j] = num
			num++
		}
		rMax--

		// left wall
		for i := rMax; i > rMin && num <= numMax; i-- {
			fmt.Println("left wall=", num)
			a[i][cMin] = num
			num++
		}

		cMin++
	}
	return a
}

func display(m, n int, a [][]int) {
	fmt.Println()
	for i := 0; i < m; i++ {
		fmt.Println()
		for j := 0; j < n; j++ {
			fmt.Printf(" %d", a[i][j])
		}
	}
	fmt.Println()
}
