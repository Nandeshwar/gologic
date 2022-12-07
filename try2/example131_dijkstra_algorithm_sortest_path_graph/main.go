package main

import (
	"container/heap"
	"fmt"
)

/*
      10    3
    0 ---- 1---|  2
  5 |          4 -- 5
    2-----3----|
       2     3

    0 to 1 weight 10 and 1 to 4 weight 3 and 4 to 5 weight 2
*/
var Graph = [][][]int{
	{{1, 10}, {2, 5}},
	{{4, 3}},
	{{3, 2}},
	{{4, 3}},
	{{5, 2}},
	{{}},
}

var tmpArr = make([]int, len(Graph))

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
	a = a[0 : l-1]
	*h = a
	return item
}

func main() {
	fmt.Println("Graph=", Graph)
	maxV := int(^uint(0) >> 1)
	for i := 1; i < len(tmpArr); i++ {
		tmpArr[i] = maxV
	}

	h := &InitHeap{}

	heap.Init(h)
	heap.Push(h, []int{0, 0})

	for h.Len() != 0 {
		// Get min weight vertex from heap
		element := heap.Pop(h)
		pair := element.([]int)
		vertex := pair[0]
		weight := pair[1]
		fmt.Println("vertex=", vertex, "weight=", weight)

		for _, pair := range Graph[vertex] {
			fmt.Println("edges=", pair)
			if len(pair) == 2 {

				adjNode := pair[0]
				edgeWeight := pair[1]

				// if weight till adjNode is less in tmpArr, update tmpArr and
				// also push adjNode with its weight from source to heap
				if weight+edgeWeight < tmpArr[adjNode] {
					tmpArr[adjNode] = weight + edgeWeight
				}
				heap.Push(h, []int{adjNode, tmpArr[adjNode]})
			}

		}
	}

	fmt.Println("tmpArr=", tmpArr)
}

/*
Graph= [[[1 10] [2 5]] [[4 3]] [[3 2]] [[4 3]] [[5 2]] [[]]]
vertex= 0 weight= 0
edges= [1 10]
edges= [2 5]
vertex= 2 weight= 5
edges= [3 2]
vertex= 3 weight= 7
edges= [4 3]
vertex= 1 weight= 10
edges= [4 3]
vertex= 4 weight= 10
edges= [5 2]
vertex= 4 weight= 10
edges= [5 2]
vertex= 5 weight= 12
edges= []
vertex= 5 weight= 12
edges= []
tmpArr= [0 10 5 7 10 12]

*/
