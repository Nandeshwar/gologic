package main

import (
	"fmt"
)

func main() {
	fmt.Println(bs1([]int{1, 2, 3, 4, 5}, 1))
}

func bs1(a []int, item int) bool {
	l := 0
	h := len(a) - 1

	for l <= h {
		mid := l + (h-l)/2
		fmt.Println(mid)

		if a[mid] == item {
			return true
		} else if item > a[mid] {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return false
}
