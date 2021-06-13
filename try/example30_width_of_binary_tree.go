// https://www.youtube.com/watch?v=Zt4bJBt2EwA
package main

import (
	"container/list"
	"fmt"
)

type Node30 struct {
	item int
	left *Node30
	right *Node30
}

func main() {
/*
			1
		0       0
	3         4   5

	ouptut: 3 without nil node consideration
	output 4 with nil node consideration
	*/
	root := &Node30 {
		item: 1,
		left: &Node30{
			item: 0,
			left: &Node30{item: 3},
		},
		right: &Node30{item: 0, left: &Node30{item:4}, right: &Node30{item: 5}},
	}

	fmt.Println(maxWidth(root))
	fmt.Println(maxWidthConsiderNullNode(root))


}
func maxWidth(root *Node30) int {
	maxWidth := 0
	size := 0 

	queue := list.New()
	queue.PushFront(root)
	

	for  {
		size = queue.Len()
		if size == 0 {
			break;
		}

		if size > maxWidth {
			maxWidth = size
		}
		
		for size != 0 {
			n := queue.Remove(queue.Front())
			node := n.(*Node30)

			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}

			size--
		}
	}
	return maxWidth
}

func maxWidthConsiderNullNode(root *Node30) int {
	maxWidth := 0
	size := 0 

	queue := list.New()
	queue.PushFront(root)
	

	for  {
		size = queue.Len()
		if size == 0 {
			break;
		}

		if size > maxWidth {
			maxWidth = size
		}
		
		for size != 0 {
			n := queue.Remove(queue.Front())
			node := n.(*Node30)

			if node != nil && node.left == nil && node.right == nil {
				size--
				break
			}
			if node != nil {
				queue.PushBack(node.left)
				queue.PushBack(node.right)
			}
			size--
		}
	}
	return maxWidth
}