package main

import (
	"fmt"
	 s "sort"
	)

func main() {
	var unSortedList = []int{ 10, 20, 30, 2, 50, 1}
	sortedList := sortFun(unSortedList)

	// User defined method
	fmt.Println(sortedList)

	// Library method
	s.Ints(unSortedList)
	fmt.Println(unSortedList)
}

func sortFun(unSortedList []int)(sortedList []int){
	for i := 0; i < len(unSortedList); i++ {
		for j := i + 1; j < len(unSortedList); j++ {
			if unSortedList[i] > unSortedList [j] {
				temp := unSortedList[i]
				unSortedList[i] = unSortedList[j]
				unSortedList[j] = temp
			}
		}
		sortedList = unSortedList
	}
	return
}
