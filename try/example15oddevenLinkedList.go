package main

import (
	"fmt"
)
type Node15 struct {
	item int
	next *Node15
}

var start *Node15
var current *Node15

func main() {
	var node *Node15
	for _, v := range []int{1, 2, 3, 4, 5} {

		node.AddItem(v)
	}
	node.PrintItem()
	node = start
	node.OddEvenLinkedList()
	fmt.Println("After odd even calculation")
	node.PrintItem()
}

func (n *Node15) OddEvenLinkedList() {
    odd := n
    even := n.next
    evenHead := even

    current = n

    for current := n; current != nil && even != nil && even.next != nil; current = current.next {
        odd.next = even.next
        odd = odd.next
        even.next = odd.next
        even = even.next
    }
    odd.next = evenHead

}


func (n *Node15) AddItem(item int) {
	if start == nil {
		n = &Node15 {
			item: item,
			next: nil,
		}
		start = n
		current = start
		return
	}

	n = current
	n.next = &Node15{
		item: item,
		next: nil,
	}
	current = n.next

}

func (n *Node15) PrintItem() {
	for node := start; node != nil; node = node.next {
		fmt.Println(node.item)		
	}
}

