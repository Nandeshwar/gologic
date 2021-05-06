package main

import "fmt"

func main() {
	numArr := []int{1, 2, 3, 4, 5}

	index := binarySearch(numArr, 5, 0, len(numArr) - 1)
	fmt.Println(index)
}

func binarySearch(numArr []int, num int, left int, right int) int {


	mid := left + (right - left) / 2

	if numArr[mid] == num {
		return mid;
	} else if num > numArr[mid] {
		left = mid + 1
	} else {
		right = mid -1 
	}

	if left > right {
		return -1
	}
	returnIem := binarySearch(numArr, num, left, right)

	return returnIem
}