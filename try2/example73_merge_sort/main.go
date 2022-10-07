package main

import (
	"fmt"
)

func main() {
	a := []int{2, 10, 3, 5, 1, 4, -1}

	a = mergeSort(a)
	fmt.Println(a)
}

func mergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}

	mid := len(a) / 2

	l1 := mergeSort(a[:mid])
	l2 := mergeSort(a[mid:])

	return merge(l1, l2)
}

func merge(l1, l2 []int) []int {

	i := 0
	j := 0

	var l []int
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			l = append(l, l1[i])
			i++
		} else {
			l = append(l, l2[j])
			j++
		}
	}

	for i < len(l1) {
		l = append(l, l1[i])
		i++
	}

	for j < len(l2) {
		l = append(l, l2[j])
		j++
	}

	return l
}
