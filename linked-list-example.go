package main

import "fmt"

type Node struct {
	item int
	next *Node
}

var (
	start *Node
	current *Node
	isFirst = true
)

func main() {
	addItem(10)
	addItem(20)
	addItem(30)
	addItem(40)
	display()
}

func addItem(item int){
	if isFirst {
		start = &Node{
			item: item,
			next: nil,
		}
		current = start
		isFirst = false
	} else {
		current.next = &Node {
			item: item,
			next: nil,
		}
		current = current.next
	}
}

func display() {
	for begin := start; ; begin = begin.next {
		fmt.Println(begin.item)
		if begin.next == nil {
			break
		}
	}
}
