package main

import (
	"fmt"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func (t *Tree) display() {
	if t == nil {
		return
	}

	fmt.Println(t.item)
	t.left.display()
	t.right.display()
}

func main() {
	t := createBinaryTree()
	t.display()
}

func createBinaryTree() *Tree {
	var input int

	fmt.Println("Enter data")
	fmt.Scanf("%d", &input)

	if input == -1 {
		return nil
	}

	t := &Tree{item: input}
	fmt.Println("Enter left for ", input)
	t.left = createBinaryTree()

	fmt.Println("Enter right for ", input)
	t.right = createBinaryTree()

	return t
}
