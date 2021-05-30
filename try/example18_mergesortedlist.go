/*
	list1 - 1, 2, 3
	list2 - 1, 2, 4
	output - 1, 1, 2, 2, 3, 4
*/
package main

import (
	"fmt"
)

type Node18 struct {
	item int
	next *Node18
}

var start18 *Node18
var current18 *Node18

func(n *Node18) createNode(item int) {
	node := &Node18 {
		item: item,
	}
	if start18 == nil {
		start18 = node
		current18 = start18
		return 
	}
	current18.next = node
	current18 = current18.next 

}

func main() {
	var n *Node18
	n.createNode(1)
	n.createNode(2)
	n.createNode(3)
	list1 := start18
	start18 = nil
	current18 = nil

	n.createNode(1)
	n.createNode(2)
	n.createNode(4)
	list2 := start18

	l := mergeList(list1, list2)

	for n := l; n != nil; n = n.next {
		fmt.Println(n.item)
	}

}


func mergeList(list1, list2 *Node18) *Node18 {
	tmp := &Node18 {
		item: -1,
	}
	start := tmp

	for list1 != nil && list2 != nil {
		if list1.item < list2.item {
			tmp.next = list1 
			list1 = list1.next
		} else {
			tmp.next = list2
			list2 = list2.next
		}

		tmp = tmp.next
	}

	if list1 != nil {
		tmp.next = list1
	} else {
		tmp.next = list2
	}
	return start.next
}

