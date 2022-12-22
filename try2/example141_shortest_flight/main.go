package main

import (
	"container/heap"
	"fmt"
)

type flight struct {
	stop   int
	vertex int
	cost   int
}

type InitHeap []flight

func (f InitHeap) Len() int           { return len(f) }
func (f InitHeap) Less(i, j int) bool { return f[i].stop < f[j].stop }
func (f InitHeap) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func (f *InitHeap) Push(item any) {
	*f = append(*f, item.(flight))
}

func (f *InitHeap) Pop() any {
	a := *f
	l := len(*f)
	item := a[l-1]
	*f = a[0 : l-1]
	return item
}

func main() {
	// time complexity: E * logV
	// if heap is replaced with Q, then time complexity will be E i.e len(flights) o(n)
	flights := [][]int{
		{0, 1, 100},
		{0, 2, 500},
		{1, 2, 100},
	}
	/*
		cost= [0 100 200]
		Min Cost= 200
	*/

	graph := make([][]flight, len(flights))

	for _, v := range flights {
		f := flight{
			stop:   1,
			vertex: v[1],
			cost:   v[2],
		}
		graph[v[0]] = append(graph[v[0]], f)
	}

	cost := findCheapestFlight(graph)
	fmt.Println("cost=", cost)
	fmt.Println("Min Cost=", cost[2])

}

func findCheapestFlight(graph [][]flight) []int {
	h := &InitHeap{}
	heap.Init(h)
	tmpArr := make([]int, len(graph))

	maxV := int(^uint(0) >> 2)
	for i := 1; i < len(tmpArr); i++ {
		tmpArr[i] = maxV
	}

	f := flight{
		stop:   0,
		vertex: 0,
		cost:   0,
	}

	heap.Push(h, f)

	for h.Len() != 0 {
		flightElement := h.Pop()
		flightItem := flightElement.(flight)

		stop := flightItem.stop
		cost := flightItem.cost
		vertex := flightItem.vertex

		stop++
		if stop > 2 {
			break
		}

		for _, v := range graph[vertex] {
			hCost := v.cost
			hVertex := v.vertex

			totalCost := cost + hCost
			if totalCost < tmpArr[hVertex] {
				tmpArr[hVertex] = totalCost
			}

			newFlight := flight{
				stop:   stop,
				vertex: hVertex,
				cost:   totalCost,
			}

			heap.Push(h, newFlight)
		}
	}
	return tmpArr
}
