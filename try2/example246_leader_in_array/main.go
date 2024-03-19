package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{17, 18, 6, 3, 4, 5}

	var leaders []int
	maxV := math.MinInt
	// Iterate from end
	// largest item in array in each index will be leaders because every item in right side will be smaller
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] > maxV {
			maxV = a[i]
			leaders = append(leaders, a[i])
		}

	}

	fmt.Println("leaders=", leaders)
}
