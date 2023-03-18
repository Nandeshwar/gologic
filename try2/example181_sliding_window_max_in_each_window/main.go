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
	"container/list"
	"fmt"
)

func main() {
	a := []int{1, 5, 2, 3, 7, 6, 8} // expected output: 5, 5, 7, 7, 8
	window := 3
	fmt.Println("original array=", a)
	m1 := maxInEachWindowAlgo1(a, window)
	fmt.Println(m1)
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

func maxInEachWindowAlgo3(a []int, window int) []int {
	var result []int
	q := list.New()

	for i := 0; i < len(a); i++ {
		// remove item from front of queue when window slides
		if q.Len() != 0 {
			frontElementIndex := q.Front().Value
			frontIndex := frontElementIndex.(int)

			if frontIndex == i-window {
				q.Remove(q.Front())
			}
		}

		// maintain max item in fron of queue:
		//in order to do this, keep removing smallest items from last of queue comparing with current
		// then insert max item index
		for q.Len() != 0 {
			lastIndexElement := q.Back().Value
			lastIndex := lastIndexElement.(int)

			if a[lastIndex] < a[i] {
				q.Remove(q.Back())
			} else {
				break
			}
		}

		// insert index
		q.PushBack(i)

		// when index >= window, add front items in result, one item from each window
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
