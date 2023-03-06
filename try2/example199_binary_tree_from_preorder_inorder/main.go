package main

import (
	"fmt"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func (t *Tree) Inorder() {
	if t == nil {
		return
	}
	t.left.Inorder()
	fmt.Println(t.item)
	t.right.Inorder()
}

func (t *Tree) Preorder() {
	if t == nil {
		return
	}
	fmt.Println(t.item)
	t.left.Preorder()
	t.right.Preorder()
}

func main() {

	preorder := []int{10, 20, 40, 50, 30, 60}
	inorder := []int{40, 20, 50, 10, 60, 30}

	// preorder = []int{3, 9, 20, 15, 7}
	// inorder = []int{9, 3, 15, 20, 7}

	root := constructTree(preorder, inorder)
	root1 := root
	root2 := root
	fmt.Println("Inorder")
	root1.Inorder()
	fmt.Println("Preorder")
	root2.Preorder()

}

func constructTree(preorder []int, inorder []int) *Tree {
	if len(preorder) <= 0 || len(inorder) <= 0 {
		return nil
	}

	m := map[int]int{}

	for i, v := range inorder {
		m[v] = i
	}

	root := &Tree{item: preorder[0]}

	mid := m[preorder[0]]

	root.left = constructTree(preorder[1:mid+1], inorder[:mid])
	root.right = constructTree(preorder[mid+1:], inorder[mid+1:])
	return root
}
