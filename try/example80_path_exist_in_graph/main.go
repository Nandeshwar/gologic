package main

type EdgesWrapper struct {
	edges []edge
}

type edge struct {
	v int // vertex
	n int // neighbour
	w int // weight
}

func checkPath(graph []EdgesWrapper, src, dest int) bool {
	visited := make([]bool, 6)
	return checkPathExists(graph, src, dest, visited)
}

func checkPathExists(graph []EdgesWrapper, src, dest int, visited []bool) bool {
	if src == dest {
		return true
	}

	visited[src] = true

	for _, e := range graph[src].edges {

		if !visited[e.n] {
			exists := checkPathExists(graph, e.n, dest, visited)
			if exists {
				return true
			}
		}
	}

	return false
}

/*
   weight of every edge is 10
  0   1   4
  2   3       5
*/
func createGraph() []EdgesWrapper {
	graph := []EdgesWrapper{
		{edges: []edge{{v: 0, n: 1, w: 10}, {0, 2, 10}}},
		{edges: []edge{{v: 1, n: 3, w: 10}, {1, 4, 10}, {1, 0, 10}}},
		{edges: []edge{{v: 2, n: 0, w: 10}, {2, 3, 10}}},
		{edges: []edge{{v: 3, n: 2, w: 10}, {3, 1, 10}}},
		{edges: []edge{{v: 4, n: 5, w: 10}, {4, 1, 10}}},
		{edges: []edge{{v: 5, n: 5, w: 10}}},
	}
	return graph
}
