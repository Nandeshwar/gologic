package main

import (
	"fmt"
)

/*
 0 --------> 1
 | ^\        |
 |        \  |
 3--------> 2
 |
 4
*/
var Graph [][]int = [][]int{
	{1, 3},
	{2},
	{0}, // if empty {}, then returns false as no cycle
	{2, 4},
}

type Status int

const (
	NotVisited Status = iota
	Visited
	InStack
)

func main() {
	fmt.Println("recursive")
	fmt.Println(dfsRecurssion(0))

}

func dfsR(v int, visited []Status) bool {
	visited[v] = InStack
	fmt.Println(v)
	if v >= len(Graph) {
		return false
	}
	for _, u := range Graph[v] {
		if visited[u] == InStack {
			return true
		}
		if visited[u] == NotVisited && dfsR(u, visited) {
			return true
		}
	}
	visited[v] = Visited
	return false
}

func dfsRecurssion(v int) bool {
	visited := make([]Status, 5)

	for i := 0; i < 5; i++ {
		if visited[i] == NotVisited && dfsR(i, visited) {
			return true
		}
	}

	return false

}
