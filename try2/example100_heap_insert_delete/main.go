package main

import (
	"fmt"
)

func main() {
	a := []int{0}
	heapInsert(&a, 50)
	heapInsert(&a, 40)
	heapInsert(&a, 45)
	heapInsert(&a, 30)
	heapInsert(&a, 20)
	heapInsert(&a, 35)
	heapInsert(&a, 10)
	heapInsert(&a, 60)
	fmt.Println(a)

	fmt.Println("Going to delete")
	fmt.Println(heapDelete(&a))
	fmt.Println(a)
}

func heapInsert(a *[]int, item int) {
	if len(*a) == 1 {
		*a = append(*a, item)
		return
	}

	n := len(*a)

	// Add item at end
	*a = append(*a, item)

	// compare item with parent and maintain max item at top
	i := n
	for i > 1 {
		parent := i / 2

		if (*a)[parent] < (*a)[i] {
			(*a)[parent], (*a)[i] = (*a)[i], (*a)[parent]
			i = parent
		} else {
			return
		}
	}
}

func heapDelete(a *[]int) int {
	n := len(*a)
	if n <= 1 {
		return 0
	}

	// in case of only 1 item
	topIndex := 1
	item := (*a)[topIndex]

	if n == 2 {
		*a = []int{(*a)[0]}
		return item
	}

	// swap last item with 1st item
	(*a)[topIndex], (*a)[n-1] = (*a)[n-1], (*a)[topIndex]

	// avoid last item
	n = n - 1
	(*a) = (*a)[0:n]

	i := 1
	for {
		leftChild := 2 * i
		rightChild := (2 * i) + 1

		// Find max of left and right child and swap top item with max of left and right
		max := 0
		// if both child present
		if leftChild < n && rightChild < n {
			if (*a)[leftChild] < (*a)[rightChild] {
				max = rightChild
			} else {
				max = leftChild
			}
		} else if leftChild < n {
			max = leftChild
		} else if rightChild < n {
			max = rightChild
		}

		// swap top item with max of left and right
		if (*a)[i] < (*a)[max] {
			(*a)[i], (*a)[max] = (*a)[max], (*a)[i]
			i = max
		} else {
			return item
		}
	}
	return item
}
