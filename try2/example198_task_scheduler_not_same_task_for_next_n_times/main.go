package main

import (
	"container/heap"
	"fmt"
)

/*
Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8

Explanation:
	A -> B -> idle -> A -> B -> idle -> A -> B
	There is at least 2 units of time between any two same tasks.


	Logic:
	  1.  map
	    A - 3
		B - 3

	  2.
	    max heap
		   3
		   3

	  3. run till 0 - n and get from heap
	     it will run 2 times since heap has 2 value
		  cycleCnt = 2


		then fill heap with
		   2
		   2

		since heap is not empty after pushing 2, 2
		    result will be n + 1
			   i.e 3

     Repeat loop:
	      cycleCnt = 2


		  heap:
		   1
		   1

		heap is not empty so
		   result will be : 3 + 3 = 6

		Repeat loop:

		cycleCnt = 2

		heap is empty so add cycleCnt to result
		  6 + 2 = 8



*/

type InitHeap []int

func (h InitHeap) Len() int           { return len(h) }
func (h InitHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h InitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Push(item any) {
	*h = append(*h, item.(int))
}

func (h *InitHeap) Pop() any {
	l := len(*h)
	item := (*h)[l-1]
	*h = (*h)[0 : l-1]
	return item
}

func main() {
	a := []string{"A", "A", "A", "B", "B", "B"}
	leastCpuUnit := findLeastCpuUnit(a, 2) // output 8

	// a = []string{"A", "A", "A", "B", "B", "B"}
	// leastCpuUnit = findLeastCpuUnit(a, 0) // output: 6

	// a = []string{"A", "A", "A", "A", "A", "A", "B", "C", "D", "E", "F", "G"}
	// leastCpuUnit = findLeastCpuUnit(a, 2) // output: 16

	fmt.Println(leastCpuUnit)
}

func findLeastCpuUnit(a []string, n int) int {
	m := make(map[string]int)
	for _, v := range a {
		m[v] += 1
	}

	fmt.Println("m=", m)

	h := &InitHeap{}
	heap.Init(h)
	for _, v := range m {
		heap.Push(h, v)
	}

	result := 0

	for h.Len() != 0 {
		var tmpArr []int
		cycleCnt := 0
		for i := 0; i <= n; i++ {
			if h.Len() != 0 {
				topElement := heap.Pop(h)
				top := topElement.(int)

				if top-1 > 0 {
					tmpArr = append(tmpArr, top-1)
				}
				cycleCnt++
			}
		}

		for _, v := range tmpArr {
			heap.Push(h, v)
		}

		if h.Len() == 0 {
			result += cycleCnt
		} else {
			result += n + 1
		}

		fmt.Println("cnt=", cycleCnt)
		fmt.Println("result=", result)
	}

	return result
}
