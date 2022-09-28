package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		{10, 20, 30},
		{40, 50, 60},
		{70, 80, 90},
	}

	cost := paintHomeWithDifferentColorMinimumCost(a)
	fmt.Println(cost)
}

func paintHomeWithDifferentColorMinimumCost(a [][]int) int {
	var i int
	for i = 1; i < 3; i++ {
		a[i][0] += min(a[i-1][1], a[i-1][2])
		a[i][1] += min(a[i-1][0], a[i-1][2])
		a[i][2] += min(a[i-1][0], a[i-1][1])
	}

	return min(min(a[i-1][0], a[i-1][1]), a[i-1][2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
