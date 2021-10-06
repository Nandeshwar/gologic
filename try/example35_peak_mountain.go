/*
	[1 2 3 2 1]
	item 3 is at peak of the mountain
	simple logic to just find max number
	Let's use binary search since items in left and right side of peaks are sorted
*/

package main

import (
	"fmt"
)
func main() {
	items := []int {1, 2, 3, 2, 1}
	indexOfPeakItem := findIndexOfPeakItem(items)
	fmt.Println(indexOfPeakItem)
}

func findIndexOfPeakItem(items []int) int {
	left := 0;
	right := len(items) 

	for left < right {
		mid := left + (right - left) / 2
		if items[mid] < items[mid + 1] {
			left = mid + 1
		} else {
			right = mid
		} 
	}
	return right;
	
}