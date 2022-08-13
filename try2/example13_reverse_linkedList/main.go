package main

import "fmt"

type Node struct {
	item int
	next *Node
}

func main() {
	three := &Node{3, nil}
	two := &Node{2, three}
	head := &Node{1, two}

	head = reverseLinkedList(head)
	for head != nil {
		fmt.Println(head.item)
		head = head.next
	}

}

// make sure prev is before head in picture
func reverseLinkedList(head *Node) *Node {
	var prev *Node = nil

	for head != nil {
		tmp := head.next
		head.next = prev
		prev = head
		head = tmp
	}
	return prev

}
