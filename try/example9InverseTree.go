package main

import (
	"fmt"
)

type TreeNode9 struct {
	Item int
	Left *TreeNode9
	Right *TreeNode9
	Children []TreeNode9
}

var root *TreeNode9

func main() {
	node4 := TreeNode9{Item: 4}
	node5 := TreeNode9{Item: 5}
	node6 := TreeNode9{Item: 6}

	node2 := TreeNode9{Item: 2, Left: &node4, Right: &node5}

	node3 := TreeNode9{Item: 3, Left: &node6}

	node1 := TreeNode9{Item: 1, Left: &node2, Right: &node3}

	root = &node1

	PrintTree(root)

	fmt.Println("After inverse ")
	root = inverseTree(root)
	PrintTree(root)

}

func inverseTree(root *TreeNode9) *TreeNode9 {
	if root == nil {
		return root
	}

	left := inverseTree(root.Left)
	right := inverseTree(root.Right)

	root.Right = left
	root.Left = right
	return root
}

func PrintTree(root *TreeNode9) {
	if root == nil {
		return
	}

	PrintTree(root.Left)
	fmt.Println(root.Item)
	PrintTree(root.Right)
}


