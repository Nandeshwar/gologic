package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	item  int
	left  *Node
	right *Node
}

/*
     10
   7   8
         9
        3
*/
func main() {
	nine := &Node{item: 9, left: &Node{item: 3}}
	seven := &Node{item: 7}
	eight := &Node{item: 8, right: nine}

	t := &Node{item: 10, left: seven, right: eight}
	fmt.Println(findMaxDepth1(t, 0))
	fmt.Println(findMaxDepth2(t))
	fmt.Println(findMaxDepth3(t))
}

func findMaxDepth1(t *Node, row int) int {
	if t == nil {
		return row
	}

	lh := findMaxDepth1(t.left, row+1)
	rh := findMaxDepth1(t.right, row+1)

	if lh > rh {
		return lh
	}
	return rh
}

func findMaxDepth2(t *Node) int {
	if t == nil {
		return 0
	}

	lh := findMaxDepth2(t.left) + 1
	rh := findMaxDepth2(t.right) + 1

	if lh > rh {
		return lh
	}
	return rh
}

func findMaxDepth3(t *Node) int {
	stack := list.New()
	stack.PushBack(t)

	height := 1
	for stack.Len() != 0 {
		rowLen := stack.Len()

		for rowLen != 0 {
			element := stack.Remove(stack.Front())
			rowLen--

			node := element.(*Node)

			if node != nil {
				if node.left != nil {
					stack.PushBack(node.left)
				}
				if node.right != nil {
					stack.PushBack(node.right)
				}
			}

			if rowLen == 0 && stack.Len() > 0 {
				height++
				rowLen = stack.Len()
			}
		}
	}
	return height
}
