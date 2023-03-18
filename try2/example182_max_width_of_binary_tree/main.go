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

/*
			10
		4		8

	  3				9

	max width = 4
*/

func main() {
	nine := &Tree{item: 9}
	eight := &Tree{item: 8, right: nine}
	three := &Tree{item: 3}
	four := &Tree{item: 4, left: three}
	root := &Tree{item: 10, left: four, right: eight}

	maxWidth := findMaxWidth(root)
	fmt.Println(maxWidth)
}

/*
	Algorithm:
	  row level traversal
	     1. root index: 0
		 2. left index: 2 * current_index + 1
		 3. right index  2 * current_index + 2
	  once row is done:
	      right_index - left_index = will be width of respective row
*/
func findMaxWidth(root *Tree) int {
	type Pair struct {
		node  *Tree
		index int
	}
	if root == nil {
		return 0
	}
	maxWidth := 1

	q := list.New()

	pair := Pair{root, 0}
	q.PushBack(pair)

	for q.Len() != 0 {
		rowLen := q.Len()
		for rowLen != 0 {
			pairElement := q.Remove(q.Front())
			rowLen--
			pairItem := pairElement.(Pair)
			node := pairItem.node
			index := pairItem.index

			leftIndex := 2*index + 1
			rightIndex := 2*index + 2

			if node.left != nil {
				pair := Pair{node.left, leftIndex}
				q.PushBack(pair)
			}
			if node.right != nil {
				pair := Pair{node.right, rightIndex}
				q.PushBack(pair)
			}
		}

		// once rowLen == 0, then lastIndex - firstIndex + 1  will be width of that row
		if q.Len() != 0 {
			firstElement := q.Front().Value
			lastElement := q.Back().Value
			firstPair := firstElement.(Pair)
			lastPair := lastElement.(Pair)

			width := lastPair.index - firstPair.index + 1
			maxWidth = max(width, maxWidth)
		}

	}

	return maxWidth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
