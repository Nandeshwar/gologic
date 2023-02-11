/*
	1. approach 1:
	     two for loop: 1 outer from 0 - n.
		     2nd for loop for window size
	     time complexity: n * k
		 space complexity: 1

	2. 1 for loop till window size
	     and maxHeap and map concept
		 time complexity: n * log(k)
		 space complexity: k

	3. 1 for loop:
	   and queue(remove from front and end)
		- add item in queue from end if previous item is less remove it
		- print it


	   time complexity: n
	   space complexitty: n
*/
package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

type InitHeap []int

func (h InitHeap) Len() int           { return len(h) }
func (h InitHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h InitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Push(item any) {
	*h = append(*h, item.(int))
}

func (h *InitHeap) Pop() any {
	a := *h
	l := len(a)
	item := a[l-1]
	a = a[:l-1]
	*h = a
	return item
}

func main() {
	a := []int{1, 5, 2, 3, 7, 6, 8} // expected output: 5, 5, 7, 7, 8
	window := 3
	fmt.Println("original array=", a)
	m1 := maxInEachWindowAlgo1(a, window)
	fmt.Println(m1)
	m2 := maxInEachWindowAlgo2(a, window)
	fmt.Println(m2)
	m3 := maxInEachWindowAlgo3(a, window)
	fmt.Println(m3)
}

func maxInEachWindowAlgo1(a []int, window int) []int {
	m := -int(uint(0) >> 2)

	var result []int
	for i := 0; i <= len(a)-window; i++ {
		for j := i; j < window+i; j++ {
			m = max(m, a[j])
		}
		result = append(result, m)
	}
	return result
}

func maxInEachWindowAlgo2(a []int, window int) []int {
	var result []int
	h := &InitHeap{}
	heap.Init(h)

	m := map[int]int{}

	for i := 0; i < len(a); i++ {
		v, ok := m[a[i]]
		if ok {
			v++
			m[a[i]] = v
		} else {
			m[a[i]] = 1
		}
	}

	for i := 0; i < window; i++ {
		heap.Push(h, a[i])
	}

	element := heap.Pop(h)
	item := element.(int)
	v := m[item]
	if v > 1 {
		v--
		m[item] = v
		heap.Push(h, item)
	} else {
		delete(m, item)
	}

	result = append(result, item)

	for i := window; i < len(a); i++ {
		heap.Push(h, a[i])

		element := heap.Pop(h)
		item := element.(int)
		v := m[item]
		if v > 1 {
			v--
			m[item] = v
			heap.Push(h, item)
		} else {
			delete(m, item)
		}

		result = append(result, item)
	}

	return result
}

func maxInEachWindowAlgo3(a []int, window int) []int {
	var result []int
	q := list.New()

	for i := 0; i < len(a); i++ {
		// remove  out of window boundary items from front - when window slides, remove 1st item
		if q.Len() != 0 {
			frontElementIndex := q.Front().Value
			frontIndex := frontElementIndex.(int)

			if frontIndex == i-window {
				q.Remove(q.Front())
			}
		}

		// keep removing smallest items from last
		for q.Len() != 0 {
			lastIndexElement := q.Back().Value
			lastIndex := lastIndexElement.(int)

			if a[lastIndex] < a[i] {
				q.Remove(q.Back())
			} else {
				break
			}
		}

		q.PushBack(i)

		// add front items in result, one item from each window
		if i >= window-1 {
			frontIndexElement := q.Front().Value
			frontIndex := frontIndexElement.(int)
			result = append(result, a[frontIndex])
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
