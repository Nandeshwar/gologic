package main

import (
	"fmt"
)

func main() {
	maxV, subArr := maxSubArr([]int{-1, 2, 10, -22, 4, 5, -1})
	fmt.Println("maxV=", maxV)
	fmt.Println("subArr=", subArr)
}

func maxSubArr(a []int) (m int, s []int) {
	maxV := 0
	subArr := []int{}

	currMaxV := 0
	currSubArr := []int{}

	for i := 0; i < len(a); i++ {
		currMaxV += a[i]
		currSubArr = append(currSubArr, a[i])

		if currMaxV < 0 {
			currMaxV = 0
			currSubArr = []int{}
		}

		if currMaxV > maxV {
			maxV = currMaxV
			subArr = currSubArr
		}

	}

	return maxV, subArr
}
