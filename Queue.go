package main

import "fmt"
var  queue []int
const queueLen = 3
var rear = -1

func main() {

	enqueue(10)
	enqueue(11)
	enqueue(12)
	enqueue(13)

	fmt.Println(queue)

	fmt.Println("Dequeued ", dequeue())
	fmt.Println("Dequeued ", dequeue())
	fmt.Println("Dequeued ", dequeue())
	fmt.Println("Dequeued ", dequeue())

	enqueue(120)
	enqueue(130)
	fmt.Println(queue)

	fmt.Println("Dequeued ", dequeue())
	fmt.Println("Dequeued ", dequeue())
}


func enqueue(item int){
	fmt.Println("length: " , len(queue))
	if len(queue) == queueLen {
		fmt.Println("Queue is full")
		return
	} else {
		rear++
		queue = append(queue, item)
		fmt.Printf("Item %d added to queue successfully\n", item)
	}
}

func dequeue() (item int){
	if len(queue) == 0 {
		fmt.Println("Queue is empty")
		return
	}

	item = queue[0]
	queue = queue[1:]
	rear--

	return item
}

