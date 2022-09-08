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
	three := &Node{item: 3}
	five := &Node{item: 5}
	four := &Node{item: 4}
	two := &Node{item: 2, left: four, right: five}
	root := &Node{item: 1, left: two, right: three}

	fmt.Println("post order: ")
	postOrder(root)
	fmt.Println("\n post order2:")
	postOrder2(root)
	fmt.Println("\nPre order: ")
	preOrder(root)
	fmt.Println("\n Preorder 2")
	preOrder2(root)
	fmt.Println("\nin order: ")
	inOrder(root)
	fmt.Println("\n in order 2")
	inOrder2(root)
	fmt.Println("\n")

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

// simple flow: get node value and add left and right in stack
// reverse the list
func postOrder2(root *Node) {
	list := list.New()

	list.PushBack(root)
	var ansList []int
	for list.Len() != 0 {
		element := list.Remove(list.Back())
		node := element.(*Node)
		ansList = append(ansList, node.item)
		if node.left != nil {
			list.PushBack(node.left)
		}
		if node.right != nil {

			list.PushBack(node.right)
		}
	}

	for i := len(ansList) - 1; i >= 0; i-- {
		fmt.Print(" ", ansList[i])
	}
}

// traverse left and put all in stack
// then take one from stack and push right node
func inOrder2(root *Node) {
	list := list.New()
	current := root
	list.PushBack(current)

	var resultList []int

	for current != nil && list.Len() != 0 {
		for current.left != nil {
			list.PushBack(current.left)
			current = current.left
		}

		element := list.Remove(list.Back())
		current = element.(*Node)
		resultList = append(resultList, current.item)
		if current.right != nil {
			list.PushBack(current.right)
			current = current.right
		}
	}

	for _, v := range resultList {
		fmt.Print(" ", v)
	}

}

// simple flow: get node value and add right and left in stack
func preOrder2(root *Node) {
	list := list.New()

	list.PushBack(root)
	var ansList []int
	for list.Len() != 0 {
		element := list.Remove(list.Back())
		node := element.(*Node)
		ansList = append(ansList, node.item)
		if node.left != nil {
			list.PushBack(node.right)
		}
		if node.right != nil {

			list.PushBack(node.left)
		}
	}

	for i := 0; i < len(ansList); i++ {
		fmt.Print(" ", ansList[i])
	}

}
