package main

import (
	"container/list"
	"fmt"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

/*
			10
		5        15
	 2    6    14   16
	1                 17

	output:
	10

5 15

2 6 14 16

1 17
Node after deletion

10

5 16

2 6 14 17

1

*/

func (t *Tree) Display() {
	queue := list.New()

	queue.PushBack(t)

	for queue.Len() != 0 {
		qLen := queue.Len()

		fmt.Println()
		for qLen != 0 {
			element := queue.Remove(queue.Front())
			node := element.(*Tree)
			qLen--

			fmt.Printf("%d ", node.item)

			if node.left != nil {
				queue.PushBack(node.left)
			}

			if node.right != nil {
				queue.PushBack(node.right)
			}
		}
		fmt.Println()
	}
}

func main() {
	seventeen := &Tree{item: 17}
	sixteen := &Tree{item: 16, right: seventeen}

	fourteen := &Tree{item: 14}
	fifteen := &Tree{item: 15, left: fourteen, right: sixteen}

	six := &Tree{item: 6}
	one := &Tree{item: 1}
	two := &Tree{item: 2, left: one}
	five := &Tree{item: 5, left: two, right: six}
	t := &Tree{item: 10, left: five, right: fifteen}
	t.Display()
	fmt.Println("Node after deletion")
	newRoot := deleteNode(t, 15)
	newRoot.Display()
}

func deleteNode(root *Tree, item int) *Tree {
	if root == nil {
		return root
	}

	if item < root.item {
		root.left = deleteNode(root.left, item)
	} else if item > root.item {
		root.right = deleteNode(root.right, item)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		root.item = minItem(root.right)
		root.right = deleteNode(root.right, root.item)

	}
	return root
}

func minItem(root *Tree) int {
	m := root.item

	for root != nil {
		m = min(m, root.item)
		root = root.left
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
