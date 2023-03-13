package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

type InitHeap []map[int]int

func (h InitHeap) Len() int { return len(h) }
func (h InitHeap) Less(i, j int) bool {
	var key1 int
	var key2 int

	for k, _ := range h[i] {
		key1 = k
	}

	for k, _ := range h[j] {
		key2 = k
	}

	return key1 < key2
}

func (h InitHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Push(item any) {
	*h = append(*h, item.(map[int]int))
}

func (h *InitHeap) Pop() any {
	a := *h
	l := len(*h)
	item := a[l-1]

	a = a[0 : l-1]

	*h = a
	return item
}

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

type Pair struct {
	horizontalDistance int
	node               *Tree
}

func main() {
	
	/*         Top View
		  |   |    |    |  |
		  |	  |	   7    |  |
		  |	  6		    2  |
		  5	    1    8  ^  3
		  ^	  4    ^    |  ^
		  |	  ^    |    |  |
		  |	  |    |    |  |
		        bottom view

			topView    = 5 6 7 2 3
			bottomView = 5 4 8 2 3
			
			
			Approach for top view: 
			   1. traverse left and right and calcuate horizontal distance -1(left) +1(right)
			      and maintain in queue: horizontal distance and node (row level traversal)
			   2. create map: key- horizontal distance: value node item value
			      if it already it exists in map, don't consider again because we traverse from left and we need top view
			   3. put the items of map in heap so that 
			   4. get the items from heap as per key and that will be top view
			
			  Note: for bottom view : point2: remove check: that check if item exists 
			          i.e keep updating 
	*/

	three := &Tree{item: 3}
	eight := &Tree{item: 8}
	two := &Tree{item: 2, left: eight, right: three}
	four := &Tree{item: 4}
	one := &Tree{item: 1, left: four}
	five := &Tree{item: 5}
	six := &Tree{item: 6, left: five, right: one}
	root := &Tree{item: 7, left: six, right: two}

	fmt.Println(topView(root))
	fmt.Println(bottomView(root))

}

func topView(root *Tree) []int {
	q := list.New()
	q.PushBack(Pair{horizontalDistance: 0, node: root})

	m := map[int]int{}

	h := &InitHeap{}
	heap.Init(h)

	for q.Len() != 0 {
		pairElement := q.Remove(q.Front())
		pair := pairElement.(Pair)

		_, ok := m[pair.horizontalDistance]
		if !ok {
			m[pair.horizontalDistance] = pair.node.item

		}

		if pair.node.left != nil {
			q.PushBack(Pair{horizontalDistance: pair.horizontalDistance - 1, node: pair.node.left})
		}

		if pair.node.right != nil {
			q.PushBack(Pair{horizontalDistance: pair.horizontalDistance + 1, node: pair.node.right})
		}
	}

	for k, v := range m {
		newMap := map[int]int{k: v}
		heap.Push(h, newMap)
	}

	var result []int
	for h.Len() != 0 {
		element := heap.Pop(h)
		m := element.(map[int]int)
		for _, v := range m {
			result = append(result, v)
		}
	}

	return result

}

func bottomView(root *Tree) []int {
	q := list.New()
	q.PushBack(Pair{horizontalDistance: 0, node: root})

	m := map[int]int{}

	h := &InitHeap{}
	heap.Init(h)

	for q.Len() != 0 {
		pairElement := q.Remove(q.Front())
		pair := pairElement.(Pair)

		m[pair.horizontalDistance] = pair.node.item

		if pair.node.left != nil {
			q.PushBack(Pair{horizontalDistance: pair.horizontalDistance - 1, node: pair.node.left})
		}

		if pair.node.right != nil {
			q.PushBack(Pair{horizontalDistance: pair.horizontalDistance + 1, node: pair.node.right})
		}
	}

	for k, v := range m {
		newMap := map[int]int{k: v}
		heap.Push(h, newMap)
	}

	var result []int
	for h.Len() != 0 {
		element := heap.Pop(h)
		m := element.(map[int]int)
		for _, v := range m {
			result = append(result, v)
		}
	}

	return result

}
