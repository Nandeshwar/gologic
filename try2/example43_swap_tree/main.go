package main

import "fmt"

type Node struct {
	item  int
	left  *Node
	right *Node
}

func main() {
	seven := &Node{item: 7}
	three := &Node{item: 3}
	two := &Node{item: 2}
	five := &Node{item: 5, left: two, right: three}
	root := &Node{item: 10, left: five, right: seven}
	root2 := &Node{item: 10, left: five, right: seven} // this is for another function

	fmt.Println("Before swaping")
	print(root)
	swap(root)
	fmt.Println("After swaping")
	print(root)

	fmt.Println("Algorithm2")
	r := swap2(root2)
	print(r)
}

func print(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.item)
	print(root.left)
	print(root.right)

}

func swap(root *Node) {
	if root == nil {
		return
	}

	swap(root.left)
	swap(root.right)

	tmp := root.left
	root.left = root.right
	root.right = tmp
}

func swap2(root *Node) *Node {
	if root == nil {
		return nil
	}

	left := swap2(root.left)
	right := swap2(root.right)

	root.left = right
	root.right = left

	return root
}
