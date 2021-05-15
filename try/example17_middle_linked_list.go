package main

import (
	"fmt"
)

type Node17 struct {
	item string;
	next *Node17
	prev *Node17
}

var head *Node17
var current *Node17

func main() {
	createNode("Sita")
	createNode("Laxman")	
	createNode("Ram")	
	createNode("Hanuman")	
	createNode("Bharat")
	
	node := middleNode()
	fmt.Println(node.item)

}

func createNode(item string) {
	
	node := Node17 {
		item: item,
	}
	if head == nil {
		head = &node
		current = head
		return
	}

	current.next = &node
	current = current.next

}

func middleNode() *Node17 {
	head1 := head;
	head2 := head.next

	for head2 != nil && head2.next != nil {
		head1= head1.next
		head2 = head2.next.next
	}
	return head1
}