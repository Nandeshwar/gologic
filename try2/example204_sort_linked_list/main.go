package main

import "fmt"

type Node struct {
	item int
	next *Node
}

func (n *Node) display() {
	for n != nil {
		fmt.Println(n.item)
		n = n.next
	}
}

func main() {
	// input: 8 3 7 4
	// output: 3 4 7 8
	four := &Node{item: 4}
	seven := &Node{item: 7, next: four}
	three := &Node{item: 3, next: seven}
	root := &Node{item: 8, next: three}

	root = sortLinkedList(root)
	root.display()
}

func sortLinkedList(root *Node) *Node {
	if root == nil || root.next == nil {
		return root
	}

	slow := root
	fast := root
	tmp := root

	for fast != nil && fast.next != nil {
		tmp = slow
		slow = slow.next
		fast = fast.next.next
	}

	tmp.next = nil

	l1 := sortLinkedList(root)
	l2 := sortLinkedList(slow)

	return merge(l1, l2)
}

func merge(l1, l2 *Node) *Node {
	tmp := &Node{item: -1}
	tmpHead := tmp

	for l1 != nil && l2 != nil {
		if l1.item < l2.item {
			tmp.next = l1
			l1 = l1.next

		} else {
			tmp.next = l2
			l2 = l2.next
		}
		tmp = tmp.next
	}

	if l1 != nil {
		tmp.next = l1
	}

	if l2 != nil {
		tmp.next = l2
	}

	return tmpHead.next

}
