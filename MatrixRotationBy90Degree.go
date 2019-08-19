package main

import (
	"fmt"
)

func main() {
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("m1: ", m1)
	m2 := rotateMatrixBy90DegreeAntiClockWise(m1)
	fmt.Println("m2: ", m2)

	fmt.Println("5 * 5")
	rotateMatrixBy90DegreeAntiClockWise2(m1)
	fmt.Println("m1: ", m1)

	m3 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	fmt.Println("m3: ", m3)
	rotateMatrixBy90DegreeAntiClockWise2(m3)
	fmt.Println("m3: ", m3)

}

// Anti-clockwise rotation
func rotateMatrixBy90DegreeAntiClockWise(m1 [][]int) (m2 [][]int) {
	m2 = make([][]int, 3)
	for i := 0; i < 3; i++ {
		m2[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m2[2-j][i] = m1[i][j]

		}
	}
	return
}

/*
Output:
m1:  [[1 2 3] [4 5 6] [7 8 9]]
m2:  [[3 6 9] [2 5 8] [1 4 7]]

*/

// clockwise rotation
func rotateMatrixBy90DegreeClockWise(m1 [][]int) (m2 [][]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m2[j][2-i] = m1[i][j]

		}
	}
	return
}

/*
Output:
m1:  [[1 2 3] [4 5 6] [7 8 9]]
m2:  [[7 4 1] [8 5 2] [9 6 3]]
*/

// Anti-clockwise rotation
func rotateMatrixBy90DegreeAntiClockWise2(m1 [][]int) {
	rowLen := len(m1)
	last := rowLen - 1
	for i := 0; i < last; i++ {
		//swapNumber(&m1[0][i], &m1[2-i][0])
		//swapNumber(&m1[0][i], &m1[2][2-i])
		//swapNumber(&m1[0][i], &m1[i][2])

		swapNumber(&m1[0][i], &m1[last-i][0])
		swapNumber(&m1[0][i], &m1[last][last-i])
		swapNumber(&m1[0][i], &m1[i][last])
	}
	return
}

func swapNumber(num1, num2 *int) {
	tmp := *num1
	*num1 = *num2
	*num2 = tmp
}

/*
m1:  [[1 2 3] [4 5 6] [7 8 9]]
m2:  [[3 6 9] [2 5 8] [1 4 7]]
5 * 5
m1:  [[3 6 9] [2 5 8] [1 4 7]]
m3:  [[1 2 3 4 5] [6 7 8 9 10] [11 12 13 14 15] [16 17 18 19 20] [21 22 23 24 25]]
m3:  [[5 10 15 20 25] [4 7 8 9 24] [3 12 13 14 23] [2 17 18 19 22] [1 6 11 16 21]]
*/
