package main

import (
	"fmt"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

/*
			1
		 1     1
	   1  1

	output: true

			1
		 1     2
	   1  1
	output: false
*/

func main() {

	one4 := &Tree{item: 1}
	one3 := &Tree{item: 1}
	one2 := &Tree{item: 1}
	one := &Tree{item: 1, left: one3, right: one4}
	root1 := &Tree{item: 1, left: one, right: one2}

	fmt.Println(checkUniValuedBinaryTree(root1))

	one4 = &Tree{item: 1}
	one3 = &Tree{item: 1}
	two := &Tree{item: 2}
	one = &Tree{item: 1, left: one3, right: one4}
	root2 := &Tree{item: 1, left: one, right: two}
	fmt.Println(checkUniValuedBinaryTree(root2))

	one4 = &Tree{item: 1}
	one3 = &Tree{item: 1}
	two = &Tree{item: 2}
	one = &Tree{item: 1, left: one3, right: one4}
	root3 := &Tree{item: 1, left: one, right: two}

	var item int
	if root3 != nil {
		item = root3.item
	}
	fmt.Println(checkUniValuedBinaryTree2(root3, item))

}

func checkUniValuedBinaryTree2(root *Tree, item int) bool {
	if root == nil {
		return true
	}

	if root.item != item {
		return false
	}

	left := checkUniValuedBinaryTree2(root.left, item)
	right := checkUniValuedBinaryTree2(root.right, item)

	return left && right
}

func checkUniValuedBinaryTree(root *Tree) bool {
	if root == nil {
		return true
	}

	if root.left != nil {
		if root.item == root.left.item {
			if root.right != nil {
				if root.item == root.right.item {
					return true
				}
				return false
			}
		}
	}

	if root.right != nil {
		if root.item == root.right.item {
			if root.left != nil {
				if root.item == root.left.item {
					return true
				}
				return false
			}
		}
	}
	// Above logic for root

	// Below call for rest of nodes
	return checkUniValuedBinaryTree(root.left) && checkUniValuedBinaryTree(root.right)
}
