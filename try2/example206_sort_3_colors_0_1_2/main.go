package main

import (
	"fmt"
)

func main() {
	colors := []int{0, 1, 2, 0, 0, 2, 2, 1, 1}
	// output: colors= [0 0 0 1 1 1 2 2 2]
	sortColors(colors)
	fmt.Println("colors=", colors)
}

/*
Algorithm:
 3 pointers
 0 - beg, mid
 1 - mid
 2 - high

 mid <= high
 when a[mid] == 0:
	swap beg, mid
	beg++
	mid++

when a[mid] == 1:
	mid++

when a[mid] == 2:
	swap mid, high
	high--
*/

func sortColors(a []int) {
	beg := 0
	mid := 0
	high := len(a) - 1

	for mid <= high {
		switch a[mid] {
		case 0:
			a[beg], a[mid] = a[mid], a[beg]
			beg++
			mid++

		case 1:
			mid++

		case 2:
			a[high], a[mid] = a[mid], a[high]
			high--
		}
	}
}
