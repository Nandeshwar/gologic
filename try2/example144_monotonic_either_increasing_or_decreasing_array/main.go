package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 1, 2, 3, 3, 4}
	fmt.Println(isMonotonicArr(a))
	a = []int{5, 4, 3, 1, 1, 0}
	fmt.Println(isMonotonicArr(a))
	a = []int{5, 4, 3, 5, 1, 1, 0}
	fmt.Println(isMonotonicArr(a))
	/*
		output:
		true
		true
		false
	*/
}

func isMonotonicArr(a []int) bool {
	ascendingOrder := true
	descendingOrder := true
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
			ascendingOrder = false
		}

		if a[i] > a[i+1] {
			descendingOrder = false
		}
	}

	return ascendingOrder || descendingOrder
}
