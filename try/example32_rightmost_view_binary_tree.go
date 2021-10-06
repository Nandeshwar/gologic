package main

import (
	"fmt"
	"container/list"
)
type Node32 struct {
	name string
	left *Node32
	right *Node32
}

func main() {

	/*
		input: RadhaKrishna
			Sita              Ram
		Radha   Krishna
		
	Laxmi Vishnu
	*/

	laxmi := &Node32{
		name: "Laxmi",
	}

	vishnu := &Node32{
		name: "Vishnu",
	}

	radha := &Node32{
		name: "Radha",
		left: laxmi,
		right: vishnu,
	}

	krishna := &Node32{
		name: "Krishna",
	}

	sita := &Node32{
		name: "Sita",
		left: radha,
		right: krishna,
	}
	ram := &Node32{
		name: "Ram",
	}

	root := &Node32{
		name: "RadhaKrishna",
		left: sita,
		right: ram,
	}

	fmt.Println(rightViewOfTree(root))
}

func rightViewOfTree(root *Node32) []string {
	var rightView []string
	queue := list.New()

	size := 1
	queue.PushBack(root)
	for size > 0 {
		
		for size > 0 {
			n := queue.Remove(queue.Front())
			node := n.(*Node32)

			size--
			if size == 0 {
				rightView = append(rightView, node.name)
			}

			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
		}

		size = queue.Len()
	}
	return rightView 
}
