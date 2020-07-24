import "fmt"

const qLen = 5

type queue [qLen]int

var front = -1
var rear = -1

func main() {
	q := queue{}
	q.enQueue(5)
	q.enQueue(10)
	q.enQueue(15)
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
	q.enQueue(50)
	fmt.Println(q.deQueue())
}

func (q *queue) enQueue(item int) {
	front++

	if front == qLen {
		fmt.Println("Queue is full")
		return
	}
	q[front] = item
}

func (q *queue) deQueue() int {
	if front == -1 {
		fmt.Println("Queue is empty")
		return -1
	}

	rear++
	item := q[rear]
	if rear == front {
		rear, front = -1, -1
	}
	return item
}
