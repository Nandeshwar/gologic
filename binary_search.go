package main

import (
	"sort"
	"fmt"
)

func main() {
	numList := []int{10, 1, 3, 4, 20, 5, 7}
	item := 20
	sort.Ints(numList)
	fmt.Println("After Sorting: ", numList)
	isFound := binarySearch(numList, item)
	fmt.Println(isFound)

	isFound = binarySearch2(numList, item, 0, len(numList) - 1)
	fmt.Println(isFound)

}

func binarySearch(numList []int, item int) (isFound bool) {

	mid := len(numList) / 2

    if mid >= len(numList)  {
    	return false
	}

	if item == numList[mid] {
		isFound = true
		return
	} else if item < numList[mid] {
		numList1 := numList[0: mid]
		return binarySearch(numList1, item )
	} else {
		numList1 := numList[mid+1:]
		return binarySearch(numList1, item )
	}
}

func binarySearch2(numList []int, item int, begin int, end int) (isFound bool){

	mid := (begin + end) / 2
	if end < begin {
		return false
	}

	if numList[mid] == item {
		return true
	} else if item < numList[mid] {
		end = mid
		return binarySearch2(numList, item, begin, end)
	} else {
		begin = mid + 1
		return binarySearch2(numList, item, begin, end)
	}
}
