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
   7     7
  3 4  4   3

*/
func main() {
	seven := &Node{item: 7, left: &Node{item: 3}, right: &Node{item: 4}}
	seven2 := &Node{item: 7, left: &Node{item: 4}, right: &Node{item: 3}}

	t := &Node{item: 10, left: seven, right: seven2}
	fmt.Println(isSymmetric(t))

}

func isSymmetric(t *Node) bool {
	t2 := t
	return isMirror(t, t2)
}

func isMirror(t1, t2 *Node) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	// if item in both tree same then continue checkingfurther
	if t1.item == t2.item {
		r1 := isMirror(t1.left, t2.right)
		r2 := isMirror(t1.right, t2.left)
		return r1 && r2
	}

	return false
}
