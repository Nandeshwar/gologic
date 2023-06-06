package main

/*
A domino has two faces
  each face contains 1 - 6

  swap domino in such a way that every domino has same value
  Input: tops = [2,1,2,4,2,2],
  bottoms =     [5,2,6,2,3,2]
	Output: 2

	- take 1st top and 1st bottom item
	- check if top or bottom do not exist in any column, then return -1

	- if bottom item in array is same as 1st top - increment


*/

import (
	"fmt"
)

func main() {
	tops := []int{2, 1, 2, 4, 2, 2}
	bottoms := []int{5, 2, 6, 2, 3, 2}
	// expectation 2

	/*
		tops = []int{3, 5, 1, 2, 3}
		bottoms = []int{3, 6, 3, 3, 4}
		// expectation -1
	*/

	top := tops[0]
	bottom := bottoms[0]

	topSwapCnt := 0
	bottomSwapCnt := 0

	for i := 0; i < len(tops); i++ {
		t := tops[i]
		b := bottoms[i]

		if top == t || top == b {

			if b == t {
				continue
			}
			if b == top {
				bottomSwapCnt++
			}
			if t == bottom {
				topSwapCnt++
			}

		} else {
			fmt.Println("Domino can not have same face")
			return
		}
	}
	if topSwapCnt == 0 {
		topSwapCnt = int(^uint(0) >> 1)
	}

	if bottomSwapCnt == 0 {
		bottomSwapCnt = int(^uint(0) >> 1)
	}

	fmt.Println(min(topSwapCnt, bottomSwapCnt))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
