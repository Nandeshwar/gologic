package main

import (
	"fmt"
)

type Node struct {
	item int
	next *Node
}

func main() {
	list1 := &Node{item: 1, next: &Node{item: 2, next: &Node{item: 3}}}
	list2 := &Node{item: 1, next: &Node{item: 4, next: &Node{item: 5}}}

	list3 := &Node{item: 0}
	list3Head := list3

	for list1 != nil && list2 != nil {
		if list1.item < list2.item {
			list3.next = list1
			list1 = list1.next
		} else {
			list3.next = list2
			list2 = list2.next
		}
		list3 = list3.next

	}

	if list1 != nil {
		list3.next = list1
		list3 = list3.next
	}

	if list2 != nil {
		list3.next = list2
		list3 = list3.next
	}

	list3Current := list3Head.next

	for list3Current != nil {
		fmt.Println(list3Current.item)
		list3Current = list3Current.next
	}
}
