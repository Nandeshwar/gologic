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
	{},
	{2, 4}, // if 2 is replaces with 0 then cycle
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

func dfsR(v int, visited []Status, parent int) bool {
	visited[v] = Visited
	fmt.Println(v)
	if v >= len(Graph) {
		return false
	}
	for _, u := range Graph[v] {
		if visited[u] == Visited && u != parent {
			return true
		}
		if visited[u] == NotVisited && dfsR(u, visited, v) {
			return true
		}
	}

	return false
}

func dfsRecurssion(v int) bool {
	visited := make([]Status, 5)

	for i, _ := range visited {
		visited[i] = NotVisited
	}

	for i := 0; i < 5; i++ {
		if visited[i] == NotVisited && dfsR(i, visited, -1) {
			return true
		}
	}

	return false

}
