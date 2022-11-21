package main

import (
	"container/heap"
	"fmt"
)

/*
Example 1:

Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Output: true
Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]
Example 2:

Input: hand = [1,2,3,4,5], groupSize = 4
Output: false
Explanation: Alice's hand can not be rearranged into groups of 4.

*/

type InitHeap []int

func (h InitHeap) Len() int           { return len(h) }
func (h InitHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h InitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Pop() any {
	a := *h
	l := len(a)
	item := a[l-1]
	a = a[:l-1]
	*h = a
	return item
}

func (h *InitHeap) Push(item any) {
	*h = append(*h, item.(int))
}

func main() {
	groupSize := 3
	a := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	//a := []int{1, 2, 2, 5, 6, 3, 6, 2, 3, 4, 7, 8} // This should fail

	l := len(a)

	if l%3 != 0 {
		fmt.Println("can not be grouped. len is not divisible by groupSize")
		return
	}

	minHeap := &InitHeap{}
	heap.Init(minHeap)

	m := map[int]int{}

	for _, v := range a {
		cnt, ok := m[v]
		if ok {
			cnt++
			m[v] = cnt
		} else {
			m[v] = 1
		}
	}

	for k, _ := range m {
		heap.Push(minHeap, k)
	}

	for minHeap.Len() != 0 {
		fmt.Println("")
		fmt.Println("minHeap=", minHeap)
		fmt.Println("map=", m)

		topE := heap.Pop(minHeap)
		top := topE.(int)
		heap.Push(minHeap, top)

		fmt.Println("top item=", top)
		fmt.Println("-------------")

		// 1. print next 3 items from map
		// 2. Keep checking next item must be incremented by 1 which must be present in map
		// 3. Reduce thier count from map if their size is 1
		// 4. Remove from minHeap if their size is 1
		endItem := top + groupSize
		for item := top; item < endItem; item++ {

			_, ok := m[item]
			if !ok {
				fmt.Println("Can not form window")
				return
			}

			fmt.Println(item)

			cnt := m[item]
			if cnt > 1 {
				cnt--
				m[item] = cnt
			} else if cnt == 1 {
				delete(m, item)
				heap.Pop(minHeap)
			}
		}

	}

}
