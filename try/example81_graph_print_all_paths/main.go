package main

import (
	"fmt"
	"strconv"
)

type EdgesWrapper struct {
	edges []edge
}

type edge struct {
	v   int // vertex
	nbr int // neighbour
	w   int // weight
}

func main() {
	visited := make([]bool, 6)
	pathVisited := "0 "
	printAllPaths(createGraph(), 0, 5, visited, pathVisited)
}

func printAllPaths(graph []EdgesWrapper, src, dest int, visited []bool, pathVisited string) {
	if src == dest {
		fmt.Println(pathVisited)
		return
	}

	visited[src] = true
	for _, e := range graph[src].edges {
		if !visited[e.nbr] {
			printAllPaths(graph, e.nbr, dest, visited, pathVisited+" "+strconv.Itoa(e.nbr))
		}
	}
	visited[src] = false

}

/*
   weight of every edge is 10
  0   1   4
  2   3       5
*/
func createGraph() []EdgesWrapper {
	graph := []EdgesWrapper{
		{edges: []edge{{v: 0, nbr: 1, w: 10}, {0, 2, 10}}},
		{edges: []edge{{v: 1, nbr: 3, w: 10}, {1, 4, 10}, {1, 0, 10}}},
		{edges: []edge{{v: 2, nbr: 0, w: 10}, {2, 3, 10}}},
		{edges: []edge{{v: 3, nbr: 2, w: 10}, {3, 1, 10}}},
		{edges: []edge{{v: 4, nbr: 5, w: 10}, {4, 1, 10}}},
		{edges: []edge{{v: 5, nbr: 4, w: 10}}},
	}
	return graph
}
