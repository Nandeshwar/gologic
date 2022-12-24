package main

import (
	"fmt"
)

func main() {
	a := []int{0, 0, 1, 1, 2, 2, 2, 3, 3, 4}
	b := make([]int, len(a))
	copy(b, a)
	c := make([]int, len(a))
	copy(c, a)

	fmt.Println("original array=", a)
	a = removeDuplicateFromSortedArr(a)
	fmt.Println(a)
	removeDuplicateFromSortedArrPtr(&b)
	fmt.Println(b)
	c = removeDuplicateFromSortedArr2(c)
	fmt.Println(c)
}

func removeDuplicateFromSortedArr(a []int) []int {
	index := 1
	for i := 0; i < len(a)-1; i++ {
		if a[i] != a[i+1] {
			a[index] = a[i+1]
			index++
		}
	}
	return a[0:index]
}

func removeDuplicateFromSortedArrPtr(a *[]int) {
	index := 1
	for i := 0; i < len(*a)-1; i++ {
		if (*a)[i] != (*a)[i+1] {
			(*a)[index] = (*a)[i+1]
			index++
		}
	}
	*a = (*a)[0:index]
}

func removeDuplicateFromSortedArr2(a []int) []int {
	index := 0
	for i := 0; i < len(a); i++ {
		j := i
		for j < len(a) && a[j] == a[i] {
			j++
		}
		a[index] = a[i]
		index++

		if j == len(a)-1 {
			a[index] = a[j]
			index++
		}
		i = j
	}
	return a[0:index]
}