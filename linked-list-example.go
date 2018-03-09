package main

import "fmt"

type Node struct {
	item int
	next *Node
	start *Node
}

var (
	current *Node
	isFirst = true
)

func main() {
	node := &Node{}
	node.addItem(10)
	node.addItem(20)
	node.addItem(30)
	node.addItem(40)
	if node.insertItem(30, 50) {
		fmt.Println("Item inserted successfully")
	} else {
		fmt.Println("Middle item not found insertion failed")
	}

	node.deleteItem(40)
	node.display()
}

func (n *Node) addItem(item int){
	if isFirst {
		n.item = item
		n.next = nil
		n.start = n
		current = n.start
		isFirst = false
	} else {
		current.next = &Node {
			item: item,
			next: nil,
		}
		current = current.next
	}
}

func (n *Node) deleteItem(item int) (isFound bool) {
	for i, current := 1, n.start; ; i, current = i + 1, current.next {
		if i == 1 && item == current.item {
			n.start = n.start.next
			return true
		}

		if current.next.item == item {
			// delete last node
			if current.next.next == nil {
				current.next = nil
				return true
			} else {
				current.next = current.next.next
				return true
			}
		}
	}
	return isFound
}

func (n *Node) insertItem(after int, item int) (isFound bool){
	for begin := n.start; ; begin = begin.next {

		if begin.item == after {
			isFound = true
			begin.next = &Node {
				item: item,
				next: begin.next,
			}
			return true
		}
		if begin.next == nil {
			break
		}
	}
	return
}

func (n *Node) display() {
	for begin := n.start; ; begin = begin.next {
		fmt.Println(begin.item)
		if begin.next == nil {
			break
		}
	}
}
