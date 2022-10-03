package main

import "fmt"

func main() {
	a := []int{5, 1, 6, 2, 4, 3}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quickSort(a []int, low, high int) {
	pivot := partition(a, low, high)

	if low < high {
		quickSort(a, low, pivot)
		quickSort(a, pivot+1, high)
	}
}

// take example: make 1st item pivot
func partition(a []int, low, high int) int {
	i := low
	j := high

	pivot := i

	for i < j {

		for a[i] <= a[pivot] {
			i++
		}

		for a[j] > a[pivot] {
			j--
		}

		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	// swap pivot with j
	a[pivot], a[j] = a[j], a[pivot]

	return j

}
