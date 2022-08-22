package main

import "fmt"

func main() {
	fmt.Println("peak number=", peakNumber([]int{1, 2, 3, 2, 1, 7, 6})) // expected 3 or 7
}

func peakNumber(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	low := 0
	high := len(arr) - 1

	for low < high {
		mid := low + (high-low)/2

		if (mid == 0 || arr[mid] > arr[mid-1]) && (mid == len(arr)-1 || arr[mid] > arr[mid+1]) {
			return mid
		}
		if arr[mid] > arr[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
