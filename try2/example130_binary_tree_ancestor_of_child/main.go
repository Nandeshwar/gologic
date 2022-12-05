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
			8		11
		7     5		   12
	*/

	five := &Tree{item: 5}
	seven := &Tree{item: 7}
	twelve := &Tree{item: 12}
	eight := &Tree{item: 8, left: seven, right: five}
	eleven := &Tree{item: 11, right: twelve}
	head := &Tree{item: 10, left: eight, right: eleven}

	//fmt.Println(findAncestor(head, 7, 5))
	fmt.Println(findAncestor(head, 11, 12))
}

func findAncestor(head *Tree, node1, node2 int) int {
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

	leftAncestor := findAncestor(head.left, node1, node2)
	if leftAncestor != 0 {
		return leftAncestor
	}
	rightAncestor := findAncestor(head.right, node1, node2)
	if rightAncestor != 0 {
		return rightAncestor
	}

	return ancestor
}
