package main

import (
	"fmt"
)

type Queue struct {
	Items []int
	Size  int
	front int
	rear  int
}

func main() {
	/*
			output:
			3
		[10 0 0]
		[10 20 0]
		[10 20 30]
		queue is full
		[10 20 30]
		10
		[-1 20 30]
		[40 20 30]
	*/
	q := New(3)
	fmt.Println(q.Size)
	q.Enqueue(10)
	fmt.Println(q.Items)
	q.Enqueue(20)
	fmt.Println(q.Items)
	q.Enqueue(30)
	fmt.Println(q.Items)
	q.Enqueue(40)
	fmt.Println(q.Items)
	fmt.Println(q.DeQueue())
	fmt.Println(q.Items)

	q.Enqueue(40)
	fmt.Println(q.Items)

}

func New(size int) *Queue {
	return &Queue{Items: make([]int, size), Size: size, front: -1, rear: -1}
}

func (q *Queue) Enqueue(item int) bool {
	if (q.rear+1)%q.Size == q.front {
		fmt.Println("queue is full")
		return false
	}

	if q.front == -1 {
		q.front = 0
	}

	q.rear = (q.rear + 1) % q.Size
	q.Items[q.rear] = item

	return true
}

func (q *Queue) DeQueue() int {
	if q.front == -1 {
		fmt.Println("queue is empty")
		return -1
	}

	item := q.Items[q.front]
	q.Items[q.front] = -1

	if q.front == q.rear {
		q.front, q.rear = -1, -1
	} else {
		q.front = (q.front + 1) % q.Size
	}
	return item
}
