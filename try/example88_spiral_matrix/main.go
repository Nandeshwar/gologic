// https://www.youtube.com/watch?v=SVFXEqn3Ceo
package main

import "fmt"

func main() {
	n := 5
	// matrix := spiralMatrix(n)
	matrix := spiralMatrixApproach2(n)

	for i := 0; i < n; i++ {
		fmt.Println()
		for j := 0; j < n; j++ {
			fmt.Print(" ", matrix[i][j])
		}
	}
	fmt.Println()
}

func spiralMatrix(n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	rMin := 0
	rMax := n

	cMin := 0
	cMax := n

	numCnt := 1

	totalMatrixLen := n * n
	for numCnt <= totalMatrixLen {
		// tranverse left wall
		for i := rMin; i < rMax && numCnt <= totalMatrixLen; i++ {
			matrix[i][cMin] = numCnt
			numCnt++
		}
		// increase left wall
		cMin++

		// traverse bottom wall
		for i := cMin; i < cMax && numCnt <= totalMatrixLen; i++ {
			matrix[rMax-1][i] = numCnt
			numCnt++
		}
		//decrease botoom wall
		rMax--

		// traverse right wall
		for i := rMax - 1; i >= rMin && numCnt <= totalMatrixLen; i-- {
			matrix[i][cMax-1] = numCnt
			numCnt++
		}

		// decrease right wall
		cMax--

		// traverse top wall
		for i := cMax - 1; i >= cMin && numCnt <= totalMatrixLen; i-- {
			matrix[rMin][i] = numCnt
			numCnt++
		}

		// increase top wall
		rMin++
	}

	return matrix
}

func spiralMatrixApproach2(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, 5)
	}

	// decide direction how you want to proceed
	// Ex: I want to go left wall first, then bottom wall, then, right wall and finally top wall
	// ok in this case
	// in case of left wall, row will change so make it 1, bottom: row will not change, right will be -1, and top will be 0
	rd := []int{1, 0, -1, 0}

	// left wall first, column will not  be changed, so 0. bottom wall - column will change so 1, then right wall no column change 0
	cd := []int{0, 1, 0, -1}

	d := 0

	r := 0
	c := 0

	maxLen := n * n
	v := 1
	for i := 0; i < maxLen; i++ {

		matrix[r][c] = v

		r += rd[d]
		c += cd[d]

		valid := validBoundary(matrix, r, c)
		if !valid {
			r -= rd[d]
			c -= cd[d]

			d = (d + 1) % 4

			r += rd[d]
			c += cd[d]

		}
		v++
	}
	return matrix

}

func validBoundary(matrix [][]int, r, c int) bool {
	if r < 0 || c < 0 || r >= len(matrix) || c >= len(matrix) || matrix[r][c] != 0 {
		return false
	}

	return true
}
