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
			8		13
		7     9	   12
	*/

	nine := &Tree{item: 5}
	seven := &Tree{item: 7}
	twelve := &Tree{item: 12}
	eight := &Tree{item: 8, left: seven, right: nine}
	thirteen := &Tree{item: 13, right: twelve}
	head := &Tree{item: 10, left: eight, right: thirteen}

	fmt.Println(findAncestor(head, 7, 9))
	fmt.Println(findAncestor(head, 13, 12))
	fmt.Println(findAncestor2(head, 13, 12))
}

func findAncestor(head *Tree, item1, item2 int) int {
	if head == nil {
		return 0
	}

	if item1 < head.item && item2 < head.item {
		return findAncestor(head.left, item1, item2)
	} else if item1 > head.item && item2 > head.item {
		return findAncestor(head.right, item1, item2)
	} else {
		return head.item
	}
}

func findAncestor2(head *Tree, node1, node2 int) int {
	var ancestor int

	if head == nil {
		return 0
	}

	if (head.left != nil && head.right != nil) && (head.left.item == node1 && head.right.item == node2) || (head.right != nil && head.right.item == node1 && head.left.item == node2) {
		return head.item
	}

	if head.item == node1 && ((head.left != nil && head.left.item == node2) || (head.right != nil && head.right.item == node2)) {
		return head.item
	}

	if head.item == node2 && ((head.left != nil && head.left.item == node1) || (head.right != nil && head.right.item == node1)) {
		return head.item
	}

	leftAncestor := findAncestor2(head.left, node1, node2)
	if leftAncestor != 0 {
		return leftAncestor
	}
	rightAncestor := findAncestor2(head.right, node1, node2)
	if rightAncestor != 0 {
		return rightAncestor
	}

	return ancestor
}
