package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type InitHeap [][]int

func (h InitHeap) Len() int           { return len(h) }
func (h InitHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h InitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Push(item any) {
	*h = append(*h, item.([]int))
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
	input := [][]int{
		{0, 30},  // 1 meeting room for this: total room=-1
		{5, 10},  // this overlaps : so another meeting room: total meeting room= 2
		{10, 12}, // this starts right after previous meeting so same room. total meeting room = 2
	}
	/*
		 1. sort by start time for loop and run loop from 2nd item
		 2. Sort by end time and put 1st item in heap
		 3. get item from heap and call it previous
		 4. start loop from 2nd item
		 5. check if current(2nd item) start date < previous end date
		     then -- overlap case
			    add current item and previous both of the items in heap (this overlap case)
			 else --- non overlap case
			    update previous end time with current(2nd end time) and put to heap

		6. count items in heap and that will be minimum meeting room count to run all the meetings


	*/

	total := findTotalMeetingRoom(input)
	fmt.Println("Total meeting room=", total)
}

func findTotalMeetingRoom(input [][]int) int {
	sort.Slice(input, func(i, j int) bool { return input[i][0] < input[j][0] })

	h := &InitHeap{}
	heap.Init(h)
	heap.Push(h, input[0])

	for i := 1; i < len(input); i++ {
		element := heap.Pop(h)
		previous := element.([]int)

		current := input[i]

		if current[0] < previous[1] {
			heap.Push(h, current)
			heap.Push(h, previous)
		} else {
			previous[1] = current[1]
			heap.Push(h, previous)
		}
	}

	return h.Len()
}
