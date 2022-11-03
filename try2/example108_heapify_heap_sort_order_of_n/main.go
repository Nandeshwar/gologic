package main

import (
	"fmt"
)

func main() {
	a := []int{0, 2, 10, 1, 5, 3}
	l := len(a) - 1
	sortWithHeap(a, l)
	fmt.Println(a)
}

func sortWithHeap(a []int, n int) {
	for ; n > 0; n-- {
		buildHeap(a, n)
		a[n], a[1] = a[1], a[n]
	}
}

func buildHeap(a []int, n int) {
	// build heap with o(n) : last insertion example is with insertion is n logn

	//n := len(a) - 1

	for i := n / 2; i > 0; i-- {
		heapify(a, n, i)
	}

	fmt.Println("heap is built=", a)
}

func heapify(a []int, n, i int) {
	larger := i

	left := 2 * i
	right := 2*i + 1

	if left <= n && a[left] > a[larger] {
		larger = left
	}

	if right <= n && a[right] > a[larger] {
		larger = right
	}

	if i != larger {
		a[i], a[larger] = a[larger], a[i]
		heapify(a, n, larger)
	}
}

/*
heap is built= [0 10 5 1 2 3]
heap is built= [0 5 3 1 2 10]
heap is built= [0 3 2 1 5 10]
heap is built= [0 2 1 3 5 10]
heap is built= [0 1 2 3 5 10]
[0 1 2 3 5 10]
*/
