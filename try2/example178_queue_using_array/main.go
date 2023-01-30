package main

import (
	"fmt"
)

type Queue struct {
	items []int
	size  int
	front int
	rear  int
}

func main() {
	q := New(3)
	fmt.Println(q.items)
	q.Enqueue(10)
	fmt.Println(q.items)
	q.Enqueue(20)
	fmt.Println(q.items)
	q.Enqueue(30)
	fmt.Println(q.items)
	q.Enqueue(40)
	fmt.Println(q.items)
	fmt.Println(q.DeQueue())
	fmt.Println(q.items)
	fmt.Println(q.DeQueue())
	fmt.Println(q.items)
	fmt.Println(q.DeQueue())
	fmt.Println(q.items)
	fmt.Println(q.DeQueue())
	fmt.Println(q.items)
	q.Enqueue(10)
	fmt.Println(q.items)
	q.Enqueue(20)
	fmt.Println(q.items)
	q.Enqueue(30)
	fmt.Println(q.items)
	q.Enqueue(40)
	fmt.Println(q.items)

}

func New(size int) *Queue {
	return &Queue{items: make([]int, size), size: size, front: -1, rear: -1}
}

func (q *Queue) Enqueue(item int) bool {
	if q.rear == q.size-1 {
		fmt.Println("Queue is full")
		return false
	}

	q.rear++
	if q.rear == 0 {
		q.front = 0
	}

	q.items[q.rear] = item
	return true
}

func (q *Queue) DeQueue() int {
	if q.front == -1 {
		fmt.Println("queue is empty")
		return -1
	}

	item := q.items[q.front]

	for i := 1; i < q.size; i++ {
		q.items[i-1] = q.items[i]
	}
	q.rear--

	if q.rear == -1 {
		q.front = -1
	}

	return item
}
