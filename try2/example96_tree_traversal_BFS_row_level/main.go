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

func main() {
	four := Tree{item: 4}
	five := Tree{item: 5}
	nine := Tree{item: 9}
	eight := Tree{item: 8, right: &nine}
	seven := Tree{item: 7, left: &four, right: &five}
	t := &Tree{item: 10, left: &seven, right: &eight}

	printEachRow(t)
}

func printEachRow(t *Tree) {
	q := list.New()
	q.PushBack(t)

	row := 0
	var rowItems []int

	for q.Len() != 0 {
		rowItemsCnt := q.Len()
		
		for rowItemsCnt > 0 {

			element := q.Remove(q.Front())
			rowItemsCnt--
			node := element.(*Tree)
			rowItems = append(rowItems, node.item)
			if rowItemsCnt == 0 {
				row++
				fmt.Println("row=", row, "items=", rowItems)
				rowItems = []int{}
			}

			if node.left != nil {
				q.PushBack(node.left)
			}
			if node.right != nil {
				q.PushBack(node.right)
			}
		}

	}
}
