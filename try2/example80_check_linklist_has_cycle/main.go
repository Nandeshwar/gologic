package main

import (
	"fmt"
)

type Node struct {
	item int
	next *Node
}

func main() {
	five := &Node{item: 5}
	four := &Node{4, five}
	three := &Node{3, four}
	two := &Node{2, three}
	five.next = two // this line causes cyclic
	s := &Node{1, two}

	fmt.Println(hasCycle(s))
}

func hasCycle(s *Node) bool {
	slow := s
	fast := s.next

	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}

		slow = slow.next
		fast = fast.next.next
	}

	return true
}
