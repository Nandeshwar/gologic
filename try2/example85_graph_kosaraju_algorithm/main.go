package main

import (
	"container/list"
	"fmt"
)

/*
 0 <--------3
 |          ^        5-----> 6 <---------7
 |          |        ^      |            |
 d          |        |    |              |
 a          |        |   d               d
 1--------> 2------> 4  a                arrow
                                         8

expected output:
   7
   8
0 3 2 1
4 6 5
*/

var Graph [][]int = [][]int{
	{1},
	{2},
	{3, 4},
	{0},
	{5},
	{6},
	{4},
	{6, 8},
	{},
}

var ReversedGraph [][]int

type Status int

const (
	NotVisited Status = iota
	Visited
	InStack
)

func main() {
	stack := list.New()
	status := make([]Status, 9)
	for i := 0; i < 9; i++ {
		if status[i] == NotVisited {
			deadEndVertexInStack(status, stack, i)
		}
	}

	ReversedGraph = reverseGraph()

	for i := 0; i < len(status); i++ {
		status[i] = NotVisited
	}

	fmt.Println("Stack Size=", stack.Len())

	// pop item from stack and do dfs() and print
	for stack.Len() != 0 {
		element := stack.Remove(stack.Back())
		v := element.(int)

		if status[v] == NotVisited {
			dfs(v, status)
			fmt.Println()
		}
	}

}

func dfs(v int, status []Status) {
	fmt.Printf("%d ", v)
	status[v] = Visited
	for _, u := range ReversedGraph[v] {
		if status[u] == NotVisited {
			dfs(u, status)
		}
	}
}

func deadEndVertexInStack(status []Status, stack *list.List, v int) {
	status[v] = Visited
	for _, u := range Graph[v] {
		if status[u] == NotVisited {
			deadEndVertexInStack(status, stack, u)
		}
	}
	stack.PushBack(v)
}

func reverseGraph() [][]int {
	graph2 := make([][]int, 9)

	for i := 0; i < len(Graph); i++ {
		for _, u := range Graph[i] {
			graph2[u] = append(graph2[u], i)
		}
	}
	return graph2
}

/*
 output:

 Stack Size= 9
7
8
0 3 2 1
4 6 5

*/
