package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	current *Node
	head    *Node
}

func main() {
	var l LinkedList
	l.createNewList(1, 2, 3, 4, 5)
	l.display()
}

func (l *LinkedList) createNewList(items ...int) {
	for i, item := range items {
		node := &Node{data: item}
		if i == 0 {
			l.head = node
			l.current = node
			continue
		}

		l.current.next = node
		l.current = node
	}
}

func (l *LinkedList) display() {
	for node := l.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}
