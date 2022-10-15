package main

import (
	"fmt"
)

func main() {
	fmt.Println(search1stBadVersion(50))
}

func search1stBadVersion(n int) int {
	beg := 1
	end := n

	min := int(^uint(0) >> 1)
	for beg <= end {
		mid := beg + (end-beg)/2

		if badVersion(mid) {
			min = Min(mid, min)
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}
	return min
}

func badVersion(n int) bool {
	if n >= 11 {
		return true
	}
	return false
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
