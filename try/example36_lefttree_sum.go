package main

import (
	"fmt"
	"container/list"
)

type Node36 struct {
	item int
	left *Node36
	right *Node36
}
func main() {
 /*
	  10
	7   20
      15 30

*/

	node10 := &Node36{
		item: 10,
		left: &Node36{
			item:7,
		},
		right: &Node36 {
			item: 20,
			left: &Node36{
				item: 15,
			},
			right: &Node36 {
				item:30,
			},
		},
	}

	fmt.Println("Sum of left nodes: ", sumOfLeftNodes(node10))

}

func sumOfLeftNodes(root *Node36) int {
	sum := 0
	q := list.New()
	q.PushFront(root)

	for q.Len() != 0 {
		rowLen := q.Len()

		for rowLen != 0 {
			e := q.Remove(q.Front())
			node := e.(*Node36)
			
			if node.right != nil {
				q.PushBack(node.right)
			}

			if node.left != nil {
				q.PushBack(node.left)
			}
			
			rowLen--;
			if rowLen == 0 {
				sum += node.item
			}
		}

		
	}
	return sum
}