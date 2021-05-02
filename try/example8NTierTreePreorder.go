package main

import (
	"fmt"
)

type TreeNode8 struct {
	Item int
	Left *TreeNode8
	Right *TreeNode8
	Children []TreeNode8
}

var root *TreeNode8

func main() {
	node4 := TreeNode8{Item: 4}
	node5 := TreeNode8{Item: 5}
	node6 := TreeNode8{Item: 6}

	node2 := TreeNode8{Item: 2, Left: &node4, Right: &node5}

	node3 := TreeNode8{Item: 3, Left: &node6}

	node1 := TreeNode8{Item: 1, Left: &node2, Right: &node3}

	root = &node1
	current := root

	//stack := Stack8{ind: -1, items: make([]TreeNode8, 10)}
	stack := Stack8{ind: -1, items: make([]interface{}, 10)}

	for current != nil || stack.ind != -1 {
		for current != nil  {
			fmt.Println(*&current.Item)
			stack.push(*current)
			current = current.Left
		}

		current = stack.pop()
		//fmt.Println(current.Item)
		current = current.Right
	}

}
type Stack8 struct {
	// items []TreeNode8
	items []interface{}
	ind int
}

func(s *Stack8) push(item TreeNode8) {
	if s.ind == 10 {
		fmt.Println("Stack Overflow")
		return
	} 

	s.ind++
	s.items[s.ind] = item
}

func(s *Stack8) pop() *TreeNode8 {
	if s.ind == -1 {
		fmt.Println("Stack is empty")
		return nil
	}
	
	item := s.items[s.ind]
	s.ind--

	nodeOfTreeNode8 := item.(TreeNode8)
	return &nodeOfTreeNode8
}

/*
 nandeshwar.sah@C02DX2V4MD6R gologic % go run try/example8NTierTreePreorder.go
1
2
4
5
3
6

*/