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

	*h = a[0 : l-1]
	return item
}

func main() {
	a := []int{2, 8, 10, 1, 5, 3, 4}
	item := findKthLargestElement(a, 2)
	fmt.Println(item)
}

func findKthLargestElement(a []int, kth int) int {
	h := &InitHeap{}
	heap.Init(h)

	cnt := 0
	for i := 0; i < len(a); i++ {
		cnt++
		heap.Push(h, a[i])

		if cnt > kth {
			heap.Pop(h)
			cnt--
		}
	}
	lastItem := heap.Pop(h)
	item := lastItem.(int)
	return item

}
