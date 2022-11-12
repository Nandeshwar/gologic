package main

import (
	"container/heap"
	"fmt"
)

type InitHeap []int

func (h InitHeap) Len() int           { return len(h) }
func (h InitHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h InitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Push(item any) {
	*h = append(*h, item.(int))
}

func (h *InitHeap) Pop() any {
	a := *h
	l := len(a)

	item := a[l-1]
	a = a[0 : l-1]
	*h = a
	return item
}
func main() {
	a := []int{2, 5, 4, 8, 6, 9}
	result := minCostRope(a)
	fmt.Println("result=", result)
}

func minCostRope(a []int) int {
	hh := InitHeap(a)
	h := &hh
	heap.Init(h)

	newRope := 0
	cost := 0
	
	// Remove two ropes and calculate cost and add the new added rope to heap
	for h.Len() >= 2 {
		element1 := heap.Pop(h)
		element2 := heap.Pop(h)

		item1 := element1.(int)
		item2 := element2.(int)

		c := item1 + item2
		cost += c

		newRope = item1 + item2
		heap.Push(h, newRope)
	}

	return cost
}
