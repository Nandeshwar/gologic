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
					5	    20
				  2   6


				output:
				found item= 6
		item = 27 is not found
	*/

	twenty := &Tree{item: 20}
	six := &Tree{item: 6}
	two := &Tree{item: 2}
	five := &Tree{item: 5, left: two, right: six}
	root := &Tree{item: 10, left: five, right: twenty}

	item := 6
	node := findBinarySearchTree(root, item)
	if node != nil {
		fmt.Println("found item=", node.item)
	} else {
		fmt.Println(fmt.Sprintf("item = %d is not found", item))
	}

	item = 27
	node = findBinarySearchTree(root, item)
	if node != nil {
		fmt.Println("found item=", node.item)
	} else {
		fmt.Println(fmt.Sprintf("item = %d is not found", item))
	}

}

func findBinarySearchTree(root *Tree, item int) *Tree {
	if root == nil {
		return nil
	}

	if item == root.item {
		return root
	}

	if item < root.item {
		return findBinarySearchTree(root.left, item)
	} else if item > root.item {
		return findBinarySearchTree(root.right, item)
	}
	return nil

}
