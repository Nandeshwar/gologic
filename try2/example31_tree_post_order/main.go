package main

import "fmt"

type Node struct {
	item  int
	left  *Node
	right *Node
}

func main() {
	three := &Node{item: 3}
	five := &Node{item: 5}
	four := &Node{item: 4}
	two := &Node{item: 2, left: four, right: five}
	root := &Node{item: 1, left: two, right: three}

	fmt.Println("post order: ")
	postOrder(root)
	fmt.Println("\nPre order: ")
	preOrder(root)
	fmt.Println("\nin order: ")
	inOrder(root)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Print(root.item, " ")
}
func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.item, " ")
	preOrder(root.left)
	preOrder(root.right)
}
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Print(root.item, " ")
	inOrder(root.right)
}
