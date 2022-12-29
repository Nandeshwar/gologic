package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	// expectation
	// 24, 12, 8, 6

	// multiply left items from index
	// left: 1 1 2 6
	// multipli right items from index and multiply with left items index value
	// right: 24, 12, 4,  1

	productOfArrayExceptOne(&a)
	fmt.Println(a)
}

func productOfArrayExceptOne(a *[]int) {
	left := make([]int, len(*a))
	right := make([]int, len(*a))

	left[0] = 1

	lp := 1
	for i := 1; i < len(*a); i++ {
		lp *= (*a)[i-1]
		left[i] = lp
	}

	right[len(*a)-1] = 1 * left[len(*a)-1]
	rp := 1
	for i := len(*a) - 2; i >= 0; i-- {
		rp *= (*a)[i+1]
		right[i] = rp * left[i]
	}
	*a = right
}
