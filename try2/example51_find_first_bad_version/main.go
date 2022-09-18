package main

import (
	"fmt"
)

func main() {
	fmt.Println(search1stBadVersion(12))
}

func search1stBadVersion(n int) int {
	beg := 1
	end := n

	for beg < end {
		mid := beg + (end-beg)/2

		if badVersion(mid) == false {
			beg = mid + 1
		} else {
			end = mid
		}
	}
	return beg
}

func badVersion(n int) bool {
	if n == 11 {
		return true
	}
	return false
}
