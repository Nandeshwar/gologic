package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchInCircularArray([]int{4, 5, 6, 1, 2, 3}, 3))
}

func searchInCircularArray(a []int, item int) int {
	beg := 0
	end := len(a) - 1

	for beg <= end {

		mid := beg + (end-beg)/2
		if a[mid] == item {
			return mid
		}

		// if true 1st portion is sorted,  else 2nd portion is sorted
		if a[beg] < a[mid] {
			// if item is withing 1st portion
			if item >= a[beg] && item < a[mid] {
				end = mid - 1
			} else {
				beg = mid + 1
			}

		} else {
			// if item is in 2nd sorted portion
			if item > a[mid] && item <= a[end] {
				beg = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
