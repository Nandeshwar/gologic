package main

import (
	"fmt"
)

/*
output:
         0
      -3    9
    -10    5
*/
type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func main() {
	a := []int{-10, -3, 0, 5, 9}
	t := arrayToBalancedBinaryTree(a, 0, len(a)-1)

	t.display()
	fmt.Println("--------")

	fmt.Println("left1=", t.left.item)
}

func arrayToBalancedBinaryTree(a []int, beg, end int) *Tree {
	t := &Tree{}

	if a == nil {
		return nil
	}

	if beg > end {
		return nil
	}

	mid := beg + (end-beg)/2

	t.item = a[mid]

	t.left = arrayToBalancedBinaryTree(a, beg, mid-1)
	t.right = arrayToBalancedBinaryTree(a, mid+1, end)

	return t
}

func (t *Tree) display() {
	if t == nil {
		return
	}

	fmt.Println(t.item)
	t.left.display()
	t.right.display()
}
