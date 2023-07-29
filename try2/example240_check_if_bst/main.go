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
		5          20
		  7      18
	*/

	// Algorithm:
	// binary search order inorder traversal always retuns items in ascending order

	eighteen := &Tree{item: 18}
	twenty := &Tree{item: 20, left: eighteen}
	seven := &Tree{item: 7}
	five := &Tree{item: 5, right: seven}
	t := &Tree{item: 10, left: five, right: twenty}

	fmt.Println(checkIFBst(t))
}

var prevItem int

func checkIFBst(t *Tree) bool {
	if t == nil {
		return true
	}

	lr := checkIFBst(t.left)
	if lr == false {
		return false
	}

	if prevItem > t.item {
		return false
	}
	prevItem = t.item

	return checkIFBst(t.right)

}
