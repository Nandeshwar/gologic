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

	fmt.Println("Recurssion ...")

	three = &Node{3, nil}
	two = &Node{2, three}
	head2 := &Node{1, two}
	head2 = reverseLinkedListRec(head2)
	for head2 != nil {
		fmt.Println(head2.item)
		head2 = head2.next
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

func reverseLinkedListRec(head *Node) *Node {

	if head == nil || head.next == nil {
		return head
	}

	newHead := reverseLinkedListRec(head.next)
	head.next.next = head
	head.next = nil

	return newHead
}
