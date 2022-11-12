package main

import (
	"container/list"
	"fmt"
)

var graph = [][]int{
	{1, 3, 4},
	{2},
	{},
	{4},
	{},
}

func main() {
	stack := list.New()
	visited := make([]bool, 5)

	for ind, _ := range graph {
		if !visited[ind] {
			topologicalOrder(graph, ind, visited, stack)
		}
	}

	for stack.Len() != 0 {
		element := stack.Remove(stack.Back())
		item := element.(int)
		fmt.Println(item)
	}

}

func topologicalOrder(graph [][]int, v int, visited []bool, stack *list.List) {
	visited[v] = true

	for _, u := range graph[v] {
		if !visited[u] {
			topologicalOrder(graph, u, visited, stack)
		}
	}

	stack.PushBack(v)
}
