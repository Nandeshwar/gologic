package main

import (
	"fmt"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func (t *Tree) display() {
	current := t

	var f func(*Tree)
	f = func(r *Tree) {
		if r == nil {
			return
		}
		fmt.Println(r.item)
		f(r.left)
		f(r.right)
	}

	f(current)
}

/*
		10
	5		20
  4  6         25
*/
func main() {

	twentyFive := &Tree{item: 25}
	twenty := &Tree{item: 20, right: twentyFive}

	six := &Tree{item: 6}
	four := &Tree{item: 4}
	five := &Tree{item: 5, left: four, right: six}
	root := &Tree{item: 10, left: five, right: twenty}

	root.display()

	root = insertInBST(root, 19)
	fmt.Println("After insertion")
	root.display()
	fmt.Println("Algo 2")
	root = insertInBST2(root, 26)
	root.display()
}

func insertInBST(root *Tree, newItem int) *Tree {
	if root == nil {
		return &Tree{item: newItem}
	}

	if newItem < root.item {
		if root.left == nil {
			root.left = &Tree{item: newItem}
		} else {
			insertInBST(root.left, newItem)
		}
	} else {
		if root.right == nil {
			root.right = &Tree{item: newItem}
		} else {
			insertInBST(root.right, newItem)
		}
	}

	return root
}

func insertInBST2(root *Tree, newItem int) *Tree {
	if root == nil {
		return &Tree{item: newItem}
	}

	if newItem < root.item {
		root.left = insertInBST2(root.left, newItem)
	} else {
		root.right = insertInBST2(root.right, newItem)
	}

	return root
}
