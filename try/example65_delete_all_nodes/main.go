/*
    To see bytes used by per operation
	go test -benchmem -run=^$ -bench .
*/
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

func main() {
	/*
			   1
		     2   3
		    4 5    7
		            8

	*/

	node8 := &Node{item: 8}
	node7 := &Node{item: 7, right: node8}

	node5 := &Node{item: 5}
	node4 := &Node{item: 4}

	node2 := &Node{item: 2, left: node4, right: node5}
	node3 := &Node{item: 3, right: node7}
	start := &Node{item: 1, left: node2, right: node3}

	deleteAllNodes(start)
	fmt.Println("Another algorithm")
	deleteAllNodesRecursive(start)
}

func deleteAllNodes(root *Node) {
	queue := list.New()
	stack := list.New()

	queue.PushBack(root)
	rowCnt := queue.Len()
	for rowCnt > 0 {

		for queue.Len() > 0 {
			obj := queue.Remove(queue.Front())
			rowCnt--
			stack.PushBack(obj)
			node := obj.(*Node)
			if node.left != nil {
				queue.PushBack(node.left)
			}

			if node.right != nil {
				queue.PushBack(node.right)
			}
		}
	}

	// for stack.Len() > 0 {
	// 	obj := stack.Remove(stack.Back())
	// 	node := obj.(*Node)
	// 	fmt.Println(node.item)
	// }
}

// traverse left and right if not nil
// if left and right nil, make current node nil
//
func deleteAllNodesRecursive(root *Node) *Node {

	if root == nil {
		return nil
	}

	if root.left != nil {
		root.left = deleteAllNodesRecursive(root.left)
	}

	if root.right != nil {
		root.right = deleteAllNodesRecursive(root.right)
	}

	if root.left == nil && root.right == nil {
		// fmt.Println(root.item)
		root = nil
		return root
	}
	return root

}
