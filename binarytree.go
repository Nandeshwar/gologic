package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BT struct {
	start   *Node
	current *Node
}

func main() {
	fmt.Println("BST")
	b := BT{}

	b.insertNode(20)
	b.insertNode(10)
	b.insertNode(5)
	b.insertNode(30)
	b.insertNode(15)

	fmt.Println("----------------")
	b.displayPreorder(b.start)

}

func (b *BT) insertNode(val int) {
	if b.current == nil {
		newNode := &Node{val: val, left: nil, right: nil}
		b.current = newNode
		if b.start == nil {
			b.start = b.current
		}
	} else {
		if val < b.current.val {
			fmt.Println("Inserted to left")
			if b.current.left != nil {
				b.current = b.current.left
			} else {
				b.current.left = &Node{val: val, left: nil, right: nil}
				return
			}
			b.insertNode(val)
		} else {
			fmt.Println("inserted to right")
			if b.current.right != nil {
				b.current = b.current.right
			} else {
				b.current.right = &Node{val: val, left: nil, right: nil}
				return
			}

			b.insertNode(val)
		}
	}
}

func (b *BT) displayPreorder(current *Node) {
	if b == nil || current == nil {
		return
	}

	fmt.Println(current.val)

	if current.left != nil {
		b.displayPreorder(current.left)
	}

	if current.right != nil {
		b.displayPreorder(current.right)
	}
}
/*
output:
BST
Inserted to left
Inserted to left
Inserted to left
inserted to right
inserted to right
Inserted to left
----------------
20
10
5
30
15
*/
