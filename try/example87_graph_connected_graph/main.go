package main

import (
	"fmt"
)

type Graph struct {
	Edges []Edge
}

type Edge struct {
	vertex    int
	neighbour int
	weight    int
}

func main() {
	// Just for fun: Beg
	a := make([][]Edge, 7)
	a[0] = []Edge{
		{0, 1, 10},
	}

	a[1] = []Edge{
		{1, 0, 10},
	}

	fmt.Println(a)

	// Just for fun: End
	visited := make([]bool, 7)
	g := createGraph()
	all_connected_graph(g, 0, visited)
}

func createGraph() []Graph {
	g := []Graph{
		Graph{Edges: []Edge{{vertex: 0, neighbour: 1, weight: 10}}},
		Graph{Edges: []Edge{{vertex: 1, neighbour: 0, weight: 10}}},
		Graph{Edges: []Edge{{vertex: 2, neighbour: 3, weight: 10}}},
		Graph{Edges: []Edge{{vertex: 3, neighbour: 2, weight: 10}}},
		Graph{Edges: []Edge{{vertex: 4, neighbour: 5, weight: 10}, Edge{vertex: 4, neighbour: 6, weight: 10}}},
		Graph{Edges: []Edge{{vertex: 5, neighbour: 6, weight: 10}, {vertex: 5, neighbour: 4, weight: 10}}},
		Graph{Edges: []Edge{{vertex: 6, neighbour: 4, weight: 10}}},
	}

	fmt.Println(g)
	return g
}

func all_connected_graph(g []Graph, index int, visited []bool) {

	for _, edges := range g {
		if !visited[index] {
			fmt.Println(index)
		}
		visited[index] = true

		for _, edge := range edges.Edges {

			if !visited[edge.neighbour] {
				all_connected_graph(g, edge.neighbour, visited)
			}
		}
	}

}

/*
output:
  [{[{0 1 10}]} {[{1 0 10}]} {[{2 3 10}]} {[{3 2 10}]} {[{4 5 10} {4 6 10}]} {[{5 6 10} {5 4 10}]} {[{6 4 10}]}]
0
1
3
2
5
6
4
bash-3.2$
*/
