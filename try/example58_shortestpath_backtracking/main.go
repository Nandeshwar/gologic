package main

// https://www.youtube.com/watch?v=S3rnLLHl0PM&list=PLKKfKV1b9e8pWy_UIiJlOlX_T4al_UtQJ&index=5

import (
	"github.com/logic-building/functional-go/fp"
)

func findShortestPath(a [][]int, x, y int) int {
	visited := make([][]bool, len(a))

	for i := 0; i < len(a); i++ {
		visited[i] = make([]bool, len(a[i]))
	}

	return shortestPath(a, visited, 0, 0, x, y)
}

func shortestPath(a [][]int, visited [][]bool, i, j, x, y int) int {
	if i < 0 || j < 0 || i >= len(a) || j >= len(a[i]) || a[i][j] != 1 || visited[i][j] {
		return 1000
	}

	if i == x && j == y {
		return 0
	}
	visited[i][j] = true

	left := shortestPath(a, visited, i, j-1, x, y) + 1
	right := shortestPath(a, visited, i, j+1, x, y) + 1
	bottom := shortestPath(a, visited, i+1, j, x, y) + 1
	top := shortestPath(a, visited, i-1, j, x, y) + 1

	visited[i][j] = false

	//return int(math.Min(float64(int(math.Min(float64(left), float64(right)))), float64(int(math.Min(float64(top), float64(bottom))))))
	return fp.MinInt([]int{left, right, bottom, top})

}
