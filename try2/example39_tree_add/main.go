package main

import "fmt"

type Node struct {
	item  int
	left  *Node
	right *Node
}

func main() {
	seven := &Node{item: 7}
	eight := &Node{item: 8}

	nine := &Node{item: 9, left: &Node{item: 3}}
	t1 := &Node{item: 10, left: seven, right: eight}
	t2 := &Node{item: 10, left: seven, right: nine}
	t3 := addTree(t1, t2)

	display(t3)

}

func addTree(t1, t2 *Node) *Node {

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	t1.item += t2.item

	t1.left = addTree(t1.left, t2.left)
	t1.right = addTree(t1.right, t2.right)

	return t1
}

func display(t3 *Node) {
	if t3 == nil {
		return
	}
	fmt.Println(t3.item)
	display(t3.left)
	display(t3.right)

}
