package main

import (
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

	fmt.Println(pathSum(t, 30))
	fmt.Println(pathSum2(t, 30))

}

func pathSum(root *Node, sum int) bool {
	if root == nil {
		return false
	}

	if sum-root.item == 0 {
		return true
	}

	if pathSum(root.left, sum-root.item) {
		return true
	}

	if pathSum(root.right, sum-root.item) {
		return true
	}
	return false
}

func pathSum2(root *Node, sum int) bool {
	if root == nil {
		return false
	}

	if root.left == nil && root.right == nil && sum-root.item == 0 {
		return true
	}

	return pathSum2(root.left, sum-root.item) || pathSum2(root.right, sum-root.item)
}
