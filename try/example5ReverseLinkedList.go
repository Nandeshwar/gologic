package main

import (
	"fmt"
)
type Node struct {
	item int
	next *Node
}

var start *Node
var current *Node

func main() {
	var node *Node
	for _, v := range []int{1, 2, 3, 4, 5} {

		node.AddItem(v)
	}
	node.PrintItem()
	prev := node.Reverse()
	fmt.Println("Reverse number.....")
	start = prev
	node.PrintItem()
}

func (n *Node) AddItem(item int) {
	if start == nil {
		n = &Node {
			item: item,
			next: nil,
		}
		start = n
		current = start
		return
	}

	n = current
	n.next = &Node{
		item: item,
		next: nil,
	}
	current = n.next

}

func (n *Node) PrintItem() {
	for node := start; node != nil; node = node.next {
		fmt.Println(node.item)		
	}
}

func(n *Node) Reverse() *Node {
	var prev *Node
	for start != nil {
		next := start.next
		start.next = prev
		prev = start
		start = next
	}

	return prev
}
