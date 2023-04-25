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
	/*
		         1
		    2        3
		4         5      6
		   7


		Burn tree from 2:
		Then result will be : 3

		Algorithm
		1. Breadth first search traversal: creating map of node and it's parent by traversing each row
		2. created visited map
		3. Traverse bfs and if any of parent or children not visited increment time count

	*/
	six := &Tree{item: 6}
	five := &Tree{item: 5}
	three := &Tree{item: 3, left: five, right: six}
	seven := &Tree{item: 7}
	four := &Tree{item: 4, left: seven}
	two := &Tree{item: 2, left: four}
	root := &Tree{item: 1, left: two, right: three}

	burnItem := 2

	m := make(map[*Tree]*Tree) // store child and parent

	targetNode := assignChildParentInMap(root, m, burnItem)

	visited := make(map[*Tree]bool)

	minCount := findMinTimeToBurnTree(targetNode, m, visited)
	fmt.Println("min count to burn tree=", minCount)

}

func assignChildParentInMap(root *Tree, m map[*Tree]*Tree, burnItem int) *Tree {
	q := list.New()
	q.PushBack(root)

	var burnItemNode *Tree

	for q.Len() != 0 {
		element := q.Remove(q.Front())
		node := element.(*Tree)

		if node.item == burnItem {
			burnItemNode = node
		}

		if node.left != nil {
			m[node.left] = node
			q.PushBack(node.left)
		}

		if node.right != nil {
			m[node.right] = node
			q.PushBack(node.right)
		}
	}
	return burnItemNode
}

func findMinTimeToBurnTree(targetNode *Tree, m map[*Tree]*Tree, visited map[*Tree]bool) int {
	burnCount := 0

	q := list.New()
	q.PushBack(targetNode)
	visited[targetNode] = true

	for q.Len() != 0 {
		flag := false
		qLen := q.Len()

		for qLen != 0 { // process each row for left, right and parent
			qLen--
			element := q.Remove(q.Front())
			node := element.(*Tree)

			_, leftNodeVisited := visited[node.left]
			if node.left != nil && !leftNodeVisited {
				flag = true
				visited[node.left] = true
				q.PushBack(node.left)
			}

			_, rightNodeVisited := visited[node.right]
			if node.right != nil && !rightNodeVisited {
				flag = true
				visited[node.right] = true
				q.PushBack(node.right)
			}

			parentNode, ok := m[node]
			if ok {
				if !visited[parentNode] {
					flag = true
					visited[parentNode] = true
					q.PushBack(parentNode)
				}
			}
		}

		if flag {
			burnCount++
		}
	}
	return burnCount
}
