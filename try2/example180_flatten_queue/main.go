package main

import (
	"container/list"
	"fmt"
)

/*
	1 -> 2 -> 3 -> 4 -> 5
	|              |
     6 -> 7-> 8   9 -> 10

    output:
     bash-3.2$ go run main.go
&{1 0xc0000a20f0 0xc0000a2090}
1
2
3
4
5
6
7
8
9
10
bash-3.2$

*/

type Node struct {
	item int
	next *Node
	down *Node
}

func main() {
	ten := &Node{item: 10}
	nine := &Node{item: 9, next: ten}
	eight := &Node{item: 8}
	seven := &Node{item: 7, next: eight}
	six := &Node{item: 6, next: seven}
	five := &Node{item: 5}
	four := &Node{item: 4, next: five, down: nine}
	three := &Node{item: 3, next: four}
	two := &Node{item: 2, next: three}
	root := &Node{item: 1, next: two, down: six}

	fmt.Println(root)
	flattenLinkedList(root)
}

func flattenLinkedList(root *Node) {
	queue := list.New()

	curr := root
	for curr != nil {
		if curr.down != nil {
			queue.PushBack(curr.down)
		}

		if curr.next == nil && queue.Len() > 0 {
			element := queue.Remove(queue.Front())
			node := element.(*Node)
			curr.next = node
		}
		curr = curr.next
	}

	curr = root
	for curr != nil {
		fmt.Println(curr.item)
		curr = curr.next
	}
}
