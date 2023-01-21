package main

/*
	if left value is missing and right value present: not complete binary tree
	otherwise: complete binary tree
*/

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
			5        7
		  2   3			9
	*/

	two := &Tree{item: 2}
	three := &Tree{item: 3}

	five := &Tree{item: 5, left: two, right: three}

	nine := &Tree{item: 9}
	seven := &Tree{item: 7, left: nil, right: nine}

	t := &Tree{
		item:  10,
		left:  five,
		right: seven,
	}

	fmt.Println(completeBinaryTree(t))
}

func completeBinaryTree(t *Tree) bool {
	if t != nil {
		if t.left == nil && t.right != nil {
			return false
		} else {
			return completeBinaryTree(t.left) && completeBinaryTree(t.right)
		}
	}
	return true
}
