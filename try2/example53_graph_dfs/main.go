package main

import (
	"container/list"
	"fmt"
)

/*
 0 --------> 1
 |           |
 |           |
 3--------> 2
 |
 4
*/
var Graph [][]int = [][]int{
	{1, 3},
	{2},
	{},
	{2, 4},
}

func main() {
	fmt.Println("recursive")
	dfsRecurssion(0)
	fmt.Println("iterative")
	dfsIterative(0)
}

func dfsR(v int, visited []bool) {
	visited[v] = true

	fmt.Println(v)

	if v >= len(Graph) {
		return
	}
	for _, u := range Graph[v] {
		if !visited[u] {
			dfsR(u, visited)
		}
	}
}

func dfsRecurssion(v int) {
	visited := make([]bool, 5)

	for i := 0; i < 5; i++ {
		if !visited[i] {
			dfsR(i, visited)
		}
	}
}

func dfsIterative(v int) {
	q := list.New()

	visited := [5]bool{}

	q.PushBack(v)

	for q.Len() != 0 {
		element := q.Remove(q.Back())
		v := element.(int)
		if !visited[v] {
			fmt.Println(v)
			visited[v] = true
			if v >= len(Graph) {
				continue
			}
			vertices := Graph[v]
			for _, v2 := range vertices {
				q.PushBack(v2)
			}
		}

	}

}
