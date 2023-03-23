package main

import (
	"fmt"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func (t *Tree) Display() {
	for t != nil {
		fmt.Println(t.item)
		t = t.right
	}
}

var head *Tree
var prev *Tree

func main() {
	six := &Tree{item: 6}
	four := &Tree{item: 4}
	one := &Tree{item: 1, left: four, right: six}
	two := &Tree{item: 2, left: one}
	five := &Tree{item: 5}
	root := &Tree{item: 3, left: five, right: two}

	convertBinaryTreeToDllInorder(root)
	
	/*
		   3
		5		2
		      1
			4   6
		output:
		5
		3
		4
		1
		6
		2
	*/

	head.Display()
}

func convertBinaryTreeToDllInorder(root *Tree) {
	if root == nil {
		return
	}

	convertBinaryTreeToDllInorder(root.left)
	if prev == nil {
		head = root
	} else {
		root.left = prev
		prev.right = root
	}

	prev = root

	convertBinaryTreeToDllInorder(root.right)

}
