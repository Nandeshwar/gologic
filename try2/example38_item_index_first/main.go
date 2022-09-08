package main

import "fmt"

func main() {
	fmt.Println(getFirstIndexOfItem([]int{1, 2, 8, 9, 10, 8, 3}, 8, 0))
}

func getFirstIndexOfItem(a []int, item int, index int) int {
	if index == len(a) {
		return -1
	}
	if a[index] == item {
		return index
	}
	index++
	return getFirstIndexOfItem(a, item, index)
}
