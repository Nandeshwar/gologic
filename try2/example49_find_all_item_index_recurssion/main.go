package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAllIndex([]int{10, 8, 3, 5, 8, 2, 8, 1}, 0, 8))
	fmt.Println(findAllIndex2([]int{10, 8, 3, 5, 8, 2, 8, 1}, 0, 8))
}

func findAllIndex(a []int, index, item int) []int {
	var arr []int
	if index == len(a) {
		return []int{}
	}

	arr = findAllIndex(a, index+1, item)

	if item == a[index] {
		arr = append(arr, index)
		return arr
	}
	return arr
}

func findAllIndex2(a []int, index, item int) []int {
	var arr []int
	if index == len(a) {
		return []int{}
	}

	if item == a[index] {
		arr = append(arr, index)
	}

	arr1 := findAllIndex(a, index+1, item)

	return append(arr, arr1...)
}
