package main

import (
	"fmt"
)

type Node struct {
	item int
	next *Node
}

func main() {
	head := &Node{item: 1, next: &Node{item: 2, next: &Node{item: 3, next: &Node{item: 4, next: &Node{item: 5}}}}}

	even := head.next // 2
	odd := head       // 1
	evenHead := even  // 2

	for odd != nil && odd.next != nil {
		odd.next = even.next
		odd = odd.next
		even.next = odd.next
		even = even.next
	}
	odd.next = evenHead

	for head != nil {
		fmt.Println(head.item)
		head = head.next
	}

}
