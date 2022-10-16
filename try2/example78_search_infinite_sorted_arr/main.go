package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	beg := 0
	end := 1

	search := 11

	for end < search {
		end = end * 2
	}

	fmt.Println("end=", end)

	index := binarySearch(a, beg, end, search)
	fmt.Println("found index=", index)

}

func binarySearch(a []int, beg, end, search int) (index int) {
	for beg <= end {
		mid := beg + (end-beg)/2

		if a[mid] == search {
			return mid
		} else if search < a[mid] {
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}
	return -1
}
