package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type InitMaxHeap []int

func (h InitMaxHeap) Len() int           { return len(h) }
func (h InitMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h InitMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InitMaxHeap) Pop() any {
	a := *h
	l := len(a)
	item := a[l-1]
	*h = a[:l-1]
	return item
}

func (h *InitMaxHeap) Push(item any) {
	*h = append(*h, item.(int))
}

type InitMinHeap []int

func (h InitMinHeap) Len() int           { return len(h) }
func (h InitMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h InitMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InitMinHeap) Pop() any {
	a := *h
	l := len(a)
	item := a[l-1]
	*h = a[:l-1]
	return item
}

func (h *InitMinHeap) Push(item any) {
	*h = append(*h, item.(int))
}

var maxHeap *InitMaxHeap
var minHeap *InitMinHeap

// for 2nd easy alorithm
var list []int

func main() {
	maxHeap = &InitMaxHeap{}
	minHeap = &InitMinHeap{}

	heap.Init(maxHeap)
	heap.Init(minHeap)

	a := []int{10, 1, 11, 12, 13, 14, 15, 16}
	for _, v := range a {
		m := findMedian(v)
		fmt.Println("median         =", m)
		fmt.Println()
	}

	fmt.Println("......Algorithm easy.......")
	for _, v := range a {
		m := findMedian2(v)
		fmt.Println("median         =", m)
		fmt.Println()
	}
}

func findMedian(item int) float64 {
	var median float64

	var maxHeapTopItem int
	if maxHeap.Len() > 0 {
		el := heap.Pop(maxHeap)
		maxHeapTopItem = el.(int)
		heap.Push(maxHeap, maxHeapTopItem)
	}
	if maxHeap.Len() == 0 || item < maxHeapTopItem {
		heap.Push(maxHeap, item)
	} else {
		heap.Push(minHeap, item)
	}

	// Balance max and min heap
	if maxHeap.Len() > minHeap.Len()+1 {
		el := heap.Pop(maxHeap)
		maxHeapTopItem := el.(int)
		heap.Push(minHeap, maxHeapTopItem)
	} else if maxHeap.Len() < minHeap.Len() {
		el := heap.Pop(minHeap)
		minHeapTopItem := el.(int)
		heap.Push(maxHeap, minHeapTopItem)
	}

	fmt.Println(maxHeap)
	fmt.Println(minHeap)

	if maxHeap.Len() == minHeap.Len() {
		maxE := heap.Pop(maxHeap)
		minE := heap.Pop(minHeap)

		maxHeapTopItem = maxE.(int)
		minHeapTopItem := minE.(int)

		heap.Push(maxHeap, maxHeapTopItem)
		heap.Push(minHeap, minHeapTopItem)

		median = float64(maxHeapTopItem+minHeapTopItem) / 2

	} else {
		el := heap.Pop(maxHeap)
		heap.Push(maxHeap, el.(int))
		median = float64(el.(int))
	}

	return median
}

func findMedian2(item int) float64 {
	var median float64
	list = append(list, item)
	sort.Ints(list)
	l := len(list)

	if l&1 != 0 {
		median = float64(list[l/2])
	} else {
		mid := l / 2
		median = float64(list[mid-1]+list[mid]) / 2
	}
	fmt.Println(list)
	return median
}

/*
bash-3.2$ go run main.go
&[10]
&[]
median         = 10

&[1]
&[10]
median         = 5.5

&[10 1]
&[11]
median         = 10

&[10 1]
&[11 12]
median         = 10.5

&[11 1 10]
&[12 13]
median         = 11

&[11 1 10]
&[12 13 14]
median         = 11.5

&[12 11 10 1]
&[13 14 15]
median         = 12

&[12 11 10 1]
&[13 14 15 16]
median         = 12.5

......Algorithm easy.......
[10]
median         = 10

[1 10]
median         = 5.5

[1 10 11]
median         = 10

[1 10 11 12]
median         = 10.5

[1 10 11 12 13]
median         = 11

[1 10 11 12 13 14]
median         = 11.5

[1 10 11 12 13 14 15]
median         = 12

[1 10 11 12 13 14 15 16]
median         = 12.5

bash-3.2$
*/
