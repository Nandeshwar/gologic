// https://www.youtube.com/watch?v=WZwNoTm_9d8&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=9
package main

import (
	"fmt"
)

type TreeNode7 struct {
	Item int
	Left *TreeNode7
	Right *TreeNode7
	Children []TreeNode7
}

var root *TreeNode7

func main() {
	node4 := TreeNode7{Item: 4}
	node5 := TreeNode7{Item: 5}
	node6 := TreeNode7{Item: 6}

	node2 := TreeNode7{Item: 2, Left: &node4, Right: &node5}

	node3 := TreeNode7{Item: 3, Left: &node6}

	node1 := TreeNode7{Item: 1, Left: &node2, Right: &node3}

	root = &node1
	current := root

	//stack := Stack7{ind: -1, items: make([]TreeNode7, 10)}
	stack := Stack7{ind: -1, items: make([]interface{}, 10)}

	for current != nil || stack.ind != -1 {
		for current != nil  {
			stack.push(*current)
			current = current.Left
		}

		current = stack.pop()
		fmt.Println(current.Item)
		current = current.Right
	}

}
type Stack7 struct {
	// items []TreeNode7
	items []interface{}
	ind int
}

func(s *Stack7) push(item TreeNode7) {
	if s.ind == 10 {
		fmt.Println("Stack Overflow")
		return
	} 

	s.ind++
	s.items[s.ind] = item
}

func(s *Stack7) pop() *TreeNode7 {
	if s.ind == -1 {
		fmt.Println("Stack is empty")
		return nil
	}
	
	item := s.items[s.ind]
	s.ind--

	nodeOfTreeNode7 := item.(TreeNode7)
	return &nodeOfTreeNode7
}

/*
 go run try/example7NTierTreeInorder.go
4
2
5
1
6
3

*/