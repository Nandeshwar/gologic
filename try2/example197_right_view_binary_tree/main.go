package main

import (
	"container/list"
	"fmt"
)

/*
			10
		5		      20
	3      6        18
	               17

	rightView: 10 20 18 17
*/

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func main() {
	seventeen := &Tree{item: 17}
	eighteen := &Tree{item: 18, left: seventeen}
	twenty := &Tree{item: 20, left: eighteen}
	six := &Tree{item: 6}
	three := &Tree{item: 3}
	five := &Tree{item: 5, left: three, right: six}
	root := &Tree{item: 10, left: five, right: twenty}

	fmt.Println(rightView1(root))
	result := []int{}
	rightView2(root, 0, &result)
	fmt.Println(result)
}

func rightView1(root *Tree) []int {
	var result []int
	queue := list.New()
	queue.PushBack(root)

	rowLen := queue.Len()
	for queue.Len() != 0 {
		element := queue.Remove(queue.Front())
		node := element.(*Tree)

		rowLen--
		if rowLen == 0 {
			result = append(result, node.item)
		}

		if node.left != nil {
			queue.PushBack(node.left)
		}
		if node.right != nil {
			queue.PushBack(node.right)
		}

		if rowLen == 0 {
			rowLen = queue.Len()
		}
	}
	return result
}

func rightView2(root *Tree, rowLevel int, result *[]int) {
	if root == nil {
		return
	}

	if rowLevel == len(*result) {
		*result = append(*result, root.item)
	}
	rightView2(root.right, rowLevel+1, result)
	rightView2(root.left, rowLevel+1, result)
}
