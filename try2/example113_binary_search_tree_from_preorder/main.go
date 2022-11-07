package main

import (
	"fmt"
)

/*
         8
       5     10
     1  7       12

   preorder: 8, 5, 1, 7, 10, 12

   construct to binary tree like above:
   so in this list : 8, 5, 1, 7, 10, 12

   fact1:
      1st itme is root: 8
   fact2:
       all items less than 8 are in left side
       all items greater than 8 are in right side

   and this is true recursively

*/

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func main() {
	a := []int{8, 5, 1, 7, 10, 12}
	t := bbt(a)
	displayTree(t)
}

func bbt(a []int) *Tree {
	if len(a) == 0 {
		return nil
	}
	if len(a) == 1 {
		return &Tree{item: a[0]}
	}

	t := Tree{item: a[0]}

	var leftIndex int
	var rightIndex int
	
	for _, item := range a[1:] {
		if item < a[0] {
			leftIndex++
		}
	}
	rightIndex = leftIndex + 1

	t.left = bbt(a[1:rightIndex])
	t.right = bbt(a[rightIndex:])
	return &t
}

func displayTree(t *Tree) {
	if t == nil {
		return
	}
	fmt.Println(t.item)
	displayTree(t.left)
	displayTree(t.right)
}
