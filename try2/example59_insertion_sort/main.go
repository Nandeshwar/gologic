package main

import (
	"fmt"
)

func main() {
	a := []int{10, 1, 11, 2, 7, 5}
	insertionSort(a)
	fmt.Println(a)
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		j := i - 1 // considered this is sorted array index
		tmp := a[i]
		for j >= 0 && a[j] > tmp {
			a[j+1] = a[j]
			j--
		}

		a[j+1] = tmp
	}
}
