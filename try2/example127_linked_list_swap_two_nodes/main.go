package main

import "fmt"

type Node struct {
	item int
	next *Node
}

func (n *Node) display() {
	head := n
	for head != nil {
		fmt.Println(head.item)
		head = head.next
	}
}

func main() {

	five := &Node{item: 5}
	four := &Node{item: 4, next: five}
	three := &Node{item: 3, next: four}
	two := &Node{item: 2, next: three}
	one := &Node{item: 1, next: two}

	one.display()
	head := SwapTwoNodes(one)
	fmt.Println("After swaping 2 nodes at a time")
	head.display()

}

func SwapTwoNodes(head *Node) *Node {
	curr := head
	tmp := &Node{item: 0}
	tmpHead := tmp

	for curr != nil && curr.next != nil {
		// 3rd node
		thirdNode := curr.next.next

		// assign 2
		tmp.next = curr.next

		// 2 points to 1
		// 2 -> 1
		tmp.next.next = curr

		// break link now: 1 points to 3 now
		curr.next = thirdNode
		curr = thirdNode

		// tmp will move to 1
		tmp = tmp.next.next
		// and then next journey
	}
	return tmpHead.next
}

/*
output:
1
2
3
4
5
After swaping 2 nodes at a time
2
1
4
3
5

*/
