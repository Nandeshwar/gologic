package main

import (
	"fmt"
)

func main() {
	// m - coloring problem
	/*
		n = 4 - vertices
		m = 3 - colors
		e = 5 - edges

		Edges[]= {(0,1), (1, 2), (2, 3), (3, 0), (0, 2)}

		0 ----- 1
		| \     |
		|    \  |
		3-------2

		check if all vertices can be colored with given m colors.
		condition no adjacent node has same color
	*/

	currentVertex := 0
	m := 3 // colors
	n := 4 // vertices

	color := make([]int, n)
	graph := [][]int{
		{1, 2, 3},
		{0, 2},
		{0, 1, 3},
		{0, 2},
	}
	// output: true

	/*
		    0         1

				2
	*/

	/*
		graph = [][]int{
			{1, 3},
			{0, 2},
			{0, 1},
		}

		m = 2
		n = 3
	*/

	// output: false

	fmt.Println("can be colored=", canBeColored(graph, m, n, currentVertex, color))
}

func canBeColored(graph [][]int, m, n, currentVertex int, color []int) bool {
	if currentVertex == n {
		return true
	}

	for currentColor := 1; currentColor <= m; currentColor++ {
		if isSafe(graph, currentVertex, currentColor, color) {
			color[currentVertex] = currentColor
			if canBeColored(graph, m, n, currentVertex+1, color) {
				return true
			}

			color[currentVertex] = 0
		}
	}

	return false
}

func isSafe(graph [][]int, currentVertex, currentColor int, color []int) bool {
	for _, neighbour := range graph[currentVertex] {
		if color[neighbour] == currentColor {
			return false
		}
	}
	return true
}
