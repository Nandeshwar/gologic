package main

import (
	"fmt"
)

func main() {
	index := findIndexFromEnd([]int{1, 2, 3, 8, 4, 5, 8, 2}, 8)
	fmt.Println("index=", index)
}

func findIndexFromEnd(a []int, item int) int {
	if len(a) == 0 {
		return -1
	}
	if item == a[len(a)-1] {
		return len(a) - 1
	}

	return findIndexFromEnd(a[0:len(a)-1], 8)
}
