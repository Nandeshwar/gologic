package main

import (
	"sort"
	"fmt"
)

func main() {
	//numList := []int{10, 1, 3, 4, 20, 5, 7}
	numList := []int{1, 2, 3, 4, 5}
	sort.Ints(numList)
	fmt.Println("After Sorting: ", numList)
	isFound := binarySearch(numList, 3)
	fmt.Println(isFound)

}

func binarySearch(numList []int, item int) (isFound bool) {

	mid := len(numList) / 2

    if mid >= len(numList){
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