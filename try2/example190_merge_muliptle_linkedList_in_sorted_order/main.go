package main

import (
	"container/heap"
	"fmt"
)

type InitHeap []*Node

func (h InitHeap) Len() int           { return len(h) }
func (h InitHeap) Less(i, j int) bool { return h[i].item < h[j].item }
func (h InitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Push(item any) {
	*h = append(*h, item.(*Node))
}

func (h *InitHeap) Pop() any {
	a := *h
	l := len(a)
	item := a[l-1]
	a = a[0 : l-1]
	*h = a
	return item
}

type Node struct {
	item int
	next *Node
}

func (n *Node) Display() {
	for n != nil {
		fmt.Println(n.item)
		n = n.next
	}
}

func main() {

	four := &Node{item: 4}
	one := &Node{item: 1, next: four}
	node10 := &Node{item: 10, next: one}
	node5 := &Node{item: 5}

	node7 := &Node{item: 7}
	node8 := &Node{item: 8, next: node7}

	nodeList := []*Node{node10, node5, node8}

	ll := mergeLinkedListInSortedOrder(nodeList)
	ll.Display()

}

func mergeLinkedListInSortedOrder(nodeList []*Node) *Node {
	h := &InitHeap{}
	heap.Init(h)

	for _, node := range nodeList {
		for node != nil {
			heap.Push(h, node)
			node = node.next
		}
	}

	tmp := &Node{item: -1}
	tmpHead := tmp

	for h.Len() != 0 {
		element := heap.Pop(h)
		item := element.(*Node)
		tmp.next = item
		tmp = tmp.next
	}
	tmp.next = nil
	return tmpHead.next
}
