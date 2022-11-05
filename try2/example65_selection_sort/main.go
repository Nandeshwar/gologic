package main

import (
	"fmt"
)

func main() {
	a := []int{1, 20, 2, 30, 3, 7, 5}
	selectionSort(a)
	fmt.Println(a)
}

func selectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if i != min {
			a[i], a[min] = a[min], a[i]
		}
	}
}
