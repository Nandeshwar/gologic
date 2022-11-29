package main

import (
	"fmt"
)

type Node struct {
	item int
	next *Node
}

/*
	1 -> 2 -> 3 -> 4 -> 5 ->
	     ^                 |
		 |-----------------
		
		output
		-------
fast and slow meet point= 5
slow node just before cycle node= 5

cycle=true, cyclePoint=2

*/

func main() {
	five := &Node{item: 5}
	four := &Node{4, five}
	three := &Node{3, four}
	two := &Node{2, three}
	five.next = two // this line causes cyclic
	s := &Node{1, two}

	cycle, item := hasCycle(s)
	fmt.Printf("\ncycle=%t, cyclePoint=%d\n", cycle, item)

	// cycleNode := findCycleNode(s)
	// fmt.Println("Cycle Node=", cycleNode.item)

	// fmt.Println("Going to remove cycle")
	// five = &Node{item: 5}
	// four = &Node{4, five}
	// three = &Node{3, four}
	// two = &Node{2, three}
	// five.next = two // this line causes cyclic
	// s = &Node{1, two}

	// h := removeCycle(s)
	// for h != nil {
	// 	fmt.Println(h.item)
	// 	h = h.next
	// }
}

func hasCycle(s *Node) (bool, int) {
	slow := s
	fast := s

	cycle := false
	for fast != nil && fast.next != nil {

		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			cycle = true
			break
		}
	}

	fmt.Println("fast and slow meet point=", slow.item)

	if cycle {
		var prev *Node
		for s != slow {
			prev = slow
			s = s.next
			slow = slow.next
		}

		fmt.Println("slow node just before cycle node=", prev.item)
	}

	return cycle, s.item
}

