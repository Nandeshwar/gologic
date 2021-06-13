package main

import (
	"fmt"
	"container/list"
)

type Node31 struct {
	item int
	left *Node31
	right *Node31
}

func main() {
/*
			1
		2       3
	3         4   5

	ouptut: 
	 1
	   2
	     3
		   3
		     4 
			   5
	*/

	root := &Node31 {
		item: 1,
		left: &Node31{
			item: 2,
			left: &Node31{item: 3},
		},
		right: &Node31{item: 3, left: &Node31{item:4}, right: &Node31{item: 5}},
	}
	printNode31(root)
	flattenBinaryTree(root)
	fmt.Println("After node flattening")
	printNode31(root)
}

func printNode31(root *Node31) {
	if root == nil {
		return
	}
	fmt.Println(root.item)
	printNode31(root.left)
	printNode31(root.right)
}

func flattenBinaryTree(root *Node31) {
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() != 0 {
		s := stack.Remove(stack.Back())
		node := s.(*Node31)

		if node.right != nil {
			stack.PushBack(node.right)
		}
		if node.left != nil {
			stack.PushBack(node.left)
		}

		var nextNode *Node31
		if stack.Len() > 0 {
			nextNodeInStackElement := stack.Back().Value
			nextNode = nextNodeInStackElement.(*Node31)
		}
		
		if node != nil {
			node.right = nextNode
			node.left = nil
		}
		
	}

}