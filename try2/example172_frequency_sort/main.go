package main

import (
	"container/heap"
	"fmt"
)

type InitHeap []map[byte]int

func (h InitHeap) Len() int { return len(h) }
func (h InitHeap) Less(i, j int) bool {
	var v1 int
	var v2 int

	for _, v := range h[i] {
		v1 = v
	}

	for _, v := range h[j] {
		v2 = v
	}

	return v1 > v2
}

func (h InitHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Pop() any {
	a := *h
	l := len(a)
	item := a[l-1]
	a = a[0 : l-1]

	*h = a
	return item
}

func (h *InitHeap) Push(item any) {
	*h = append(*h, item.(map[byte]int))
}

func main() {
	word := "tree"
	// ouput: sortedWord= eetr
	sortedWord := sortByFrequency(word)
	fmt.Println("sortedWord=", sortedWord)
}

func sortByFrequency(word string) string {
	h := &InitHeap{}
	heap.Init(h)

	var sortedSort []byte
	m := map[byte]int{}
	for _, v := range word {
		m[byte(v)] = m[byte(v)] + 1
	}

	//convert map to array of pair
	for k, v := range m {
		heap.Push(h, map[byte]int{k: v})
	}

	for h.Len() != 0 {
		element := heap.Pop(h)
		m := element.(map[byte]int)

		for k, v := range m {
			for i := 0; i < v; i++ {
				sortedSort = append(sortedSort, k)
			}
		}

	}

	return string(sortedSort)
}
