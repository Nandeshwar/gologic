package main

import (
	"fmt"
)

func main() {
	a := []int{10, 5, 7, 3, 8, 1, 2}
	bubbleSort(a)
	fmt.Println(a)
}

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
