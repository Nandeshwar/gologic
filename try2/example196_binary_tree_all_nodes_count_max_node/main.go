package main

import (
	"fmt"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func main() {
	/*
					10
				  5     17
				3  4      18
				            19

		totalNodes= 7
		max node= 19

	*/

	nineteen := &Tree{item: 19}
	eighteen := &Tree{item: 18, right: nineteen}
	seventeen := &Tree{item: 17, right: eighteen}
	four := &Tree{item: 4}
	three := &Tree{item: 3}
	five := &Tree{item: 5, left: three, right: four}
	t := &Tree{item: 10, left: five, right: seventeen}

	totalNodes := TotalNodes(t)
	fmt.Println("totalNodes=", totalNodes)
	fmt.Println("max node=", maxNode(t))
}

func TotalNodes(t *Tree) int {
	if t == nil {
		return 0
	}

	// rest node from root then + 1(for root)
	return TotalNodes(t.left) + TotalNodes(t.right) + 1
}

func maxNode(t *Tree) int {

	if t == nil {
		return -1
	}

	return max(t.item, max(maxNode(t.left), maxNode(t.right)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
