package main

import "fmt"

type Node struct {
	item  int
	left  *Node
	right *Node
}

func main() {

	seven := &Node{item: 7}
	four := &Node{item: 4}
	three := &Node{item: 3}
	five := &Node{item: 5, left: three, right: four}
	t := &Node{item: 10, left: five, right: seven}

	s := &Node{item: 5, left: three, right: four}

	fmt.Println(isSubTree(s, t))

}

func isSubTree(s, t *Node) bool {
	if t == nil {
		return false
	}
	if isSameTree(s, t) {
		return true
	}

	return isSubTree(s, t.left) || isSubTree(s, t.right)

}

func isSameTree(s, t *Node) bool {
	if s == nil || t == nil {
		return s == nil && t == nil
	}
	if s.item == t.item {
		return isSameTree(s.left, t.left) && isSameTree(s.right, t.right)
	}
	return false
}
