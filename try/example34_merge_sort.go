// nlog(n)
package main

import (
	"fmt"
)
func main() {
	list := []int{20, 10, 2, 1, 3, 4}

	sortedList := mergeSort(list)
	fmt.Println(sortedList)
}

func mergeSort(list []int) []int {

	if len(list) == 1 {
		return list
	}
	mid := len(list) / 2
    list1 := make([]int, mid)
	list2 := make([]int, len(list) - mid) 

	j := 0
	k := 0
	for i, v := range list {
		if i < mid {
			list1[j] = v
			j++
		} else {
			list2[k] = v
			k++
		}
	}

	l1 := mergeSort(list1)
	l2 := mergeSort(list2)


	return merge(l1, l2)
}

func merge(list1, list2 []int) []int {
	newList := make([]int, len(list1) + len(list2))

	i := 0;
	j := 0;
	k :=0;

	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			newList[k] = list1[i]
			i++
		} else {
			newList[k] = list2[j]
			j++
		}
		k++
	}

	for i < len(list1) {
		newList[k] = list1[i]
		i++
		k++ 
	}

	for j < len(list2) {
		newList[k] = list2[j]
		j++
		k++ 
	}

	return newList
	
}