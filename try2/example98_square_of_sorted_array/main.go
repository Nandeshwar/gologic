package main

import (
	"fmt"
)

func main() {
	fmt.Println(squareOfSortedArr([]int{-4, -1, 0, 3, 5}))
}

// find 1st negative index and 1st positive index
// then compare squre of 1st negative index value and 1st positive index value
// and store smallest number in new array and increment or decrement respective pointer

func squareOfSortedArr(a []int) []int {
	b := make([]int, len(a))

	negativeNumberIndex := -1
	positiveNumberIndex := -1
	for ind, v := range a {
		if v >= 0 {
			negativeNumberIndex = ind - 1
			positiveNumberIndex = ind
			break
		}
	}

	index := 0
	for negativeNumberIndex >= 0 && positiveNumberIndex < len(a) {
		ns := a[negativeNumberIndex] * a[negativeNumberIndex]
		ps := a[positiveNumberIndex] * a[positiveNumberIndex]

		if ns < ps {
			b[index] = ns
			negativeNumberIndex--
		} else {
			b[index] = ps
			positiveNumberIndex++
		}
		index++
	}

	for negativeNumberIndex >= 0 {
		ns := a[negativeNumberIndex] * a[negativeNumberIndex]
		negativeNumberIndex--
		b[index] = ns
		index++
	}

	for positiveNumberIndex < len(a) {
		ps := a[positiveNumberIndex] * a[positiveNumberIndex]
		positiveNumberIndex++
		b[index] = ps
		index++
	}
	return b
}
