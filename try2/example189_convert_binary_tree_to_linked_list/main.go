package main

import (
	"container/list"
	"fmt"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func (t *Tree) Display() {
	current := t
	for current != nil {
		fmt.Println(current.item)
		current = current.right
	}
}
func main() {
	/*
				10
			5		   15

			         11
		3      4

		output:
		--------
		10
		  5
		   3
		    4
			 15
			   11
	*/
	three := &Tree{item: 3}
	four := &Tree{item: 4}
	eleven := &Tree{item: 11}
	five := &Tree{item: 5, left: three, right: four}
	fifteen := &Tree{item: 15, left: eleven}
	t := &Tree{item: 10, left: five, right: fifteen}
	ll := convertTreeToLinkedList(t)
	ll.Display()
}

func convertTreeToLinkedList(t *Tree) *Tree {
	if t == nil {
		return nil
	}

	stack := list.New()

	stack.PushBack(t)

	for stack.Len() != 0 {
		element := stack.Remove(stack.Back())
		node := element.(*Tree)

		if node.right != nil {
			stack.PushBack(node.right)
		}

		if node.left != nil {
			stack.PushBack(node.left)
		}

		if stack.Len() != 0 {

			topElement := stack.Back().Value
			topNode := topElement.(*Tree)
			node.right = topNode
			node.left = nil
		}
	}
	return t

}
