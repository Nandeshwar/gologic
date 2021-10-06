// https://www.youtube.com/watch?v=j16cwbLEf9w&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=42
package main

import "fmt"
import "container/list"


type Node29 struct {
	item int
	left *Node29
	right *Node29
}

func main() {
/*
			1
		0       0
	3

	
	*/
	root := &Node29 {
		item: 1,
		left: &Node29{
			item: 0,
			left: &Node29{item: 3},
		},
		right: &Node29{item: 0},
	}

	ans := true
	completenessOfBinaryTree(root, &ans)
	fmt.Println(ans)

	// --------------------Another approach -------------------------------
	queue := list.New()
	queue.PushFront(root)

	foundNil := false

	ans = true;
	for queue.Len() != 0 {
		n := queue.Remove(queue.Front())
		node := n.(*Node29)
		if node == nil {
			foundNil = true
		} else {
			if foundNil {
				ans = false;
			}
			queue.PushBack(node.left)
			queue.PushBack(node.right)
		}
	}

	fmt.Println("Answer: ", ans)
}

func completenessOfBinaryTree(root *Node29, ans *bool) bool {
	if root == nil {
		return false
	}

	if (root.left == nil && root.right != nil) {
		*ans = false

		return false
	}

	left := completenessOfBinaryTree(root.left, ans)
	 completenessOfBinaryTree(root.right, ans)

	return left

}