package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type InitHeap [][]int

func (h InitHeap) Len() int { return len(h) }
func (h InitHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1]
}
func (h InitHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Push(item any) {
	v := item.([]int)
	*h = append(*h, v)
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
	a := [][]int{
		{1, 3},
		{-2, 2},
	}

	h := InitHeap(a)
	sort.Sort(h)
	fmt.Println("sorted h=", h)

	h2 := &InitHeap{}
	heap.Init(h2)

	cnt := 0
	for _, v := range a {
		cnt++
		heap.Push(h2, v)
		if cnt == 2 {
			heap.Pop(h2)
		}
	}

	element := heap.Pop(h2)
	item := element.([]int)
	fmt.Println("result=", item)
}
