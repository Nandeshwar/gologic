// https://www.youtube.com/watch?v=77LJc56bwnE&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=34
package main

import (
	"fmt"
)

type Node24 struct {
	item int
	left *Node24
	right *Node24
}

func main() {
/*
	   1
   0      0
        0   1

	ouput:
	  1
	    0
		   1
*/

	root := Node24{
		item: 1,
		left: &Node24{item: 0},
		right: &Node24{
			item: 0, 
			left: &Node24{item:0}, 
			right: &Node24{item:1},
		},
	}

	print(&root)
	pruneNode(&root)
	fmt.Println("After prunning")
	print(&root)
	
}

func pruneNode(node *Node24) bool {
	if node == nil {
		return false
	}
	leftHas1 := pruneNode(node.left)
	rightHas1 := pruneNode(node.right)

	if !leftHas1 {
		node.left = nil
	}
	if !rightHas1 {
		node.right = nil
	}

	if node.item == 1 {
		return true
	} else {
		if leftHas1 || rightHas1 {
			return true
		} else {
			return false
		}
	}
}


func print(root *Node24) {
	if root == nil {
		return
	}
	fmt.Println(root.item)
	print(root.left)
	print(root.right)
}