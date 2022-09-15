package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{1, 2, 8}))
}

func plusOne(a []int) []int {
	var i int
	for i = len(a) - 1; i >= 0; i-- {
		if a[i] < 9 {
			a[i] = a[i] + 1
			return a
		}
		if a[i] == 9 {
			a[i] = 0
		}
	}
	var b []int
	if i == -1 {
		b = append(b, 1)
		b = append(b, a...)
	}
	return b
}
