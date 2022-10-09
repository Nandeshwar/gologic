package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[j] < h[i] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	// 1, 7, 2, 3, 8, 1, 1 : collide two max stones : 7, 8 here remaining 1 . 1 will added
	// 1, 2, 3, 1, 1, 1 : collide 3 and 2 and remaining 1 will be added
	// 1, 1, 1, 1, 1 : collide 1, 1 nothing will be added
	// 1 1 1 : collide 1 and 1 nothing will be added
	// 1 will be answer
	a := []int{1, 7, 2, 3, 8, 1, 1}

	h := &IntHeap{}
	heap.Init(h)
	for _, item := range a {
		heap.Push(h, item)
	}
	// fmt.Printf("minimum: %d\n", (*h)[0])
	// for h.Len() > 0 {
	// 	fmt.Printf("%d ", heap.Pop(h))
	// }

	for h.Len() >= 2 {
		element1 := heap.Pop(h)
		element2 := heap.Pop(h)
		item1 := element1.(int)
		item2 := element2.(int)

		fmt.Println("item1 popped=", item1)
		fmt.Println("item2 popped=", item2)

		if item1 > item2 {
			heap.Push(h, item1-item2)
		}
	}
	if h.Len() > 0 {
		fmt.Println("answer=", heap.Pop(h))
	} else {
		fmt.Println("answer is empty")
	}

}
