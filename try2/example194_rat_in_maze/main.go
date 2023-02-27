package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 1},
	}

	m := len(a)
	n := len(a[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	paths := ratInMaze(0, 0, m, n, a, visited, "")
	fmt.Println("paths=", paths)

	/*
		output:
		paths= [DDRR DRDR]
	*/
}

func ratInMaze(i, j, m, n int, a [][]int, visited [][]bool, path string) []string {
	var paths []string

	if i < 0 || j < 0 || i >= m || j >= n || visited[i][j] == true || a[i][j] == 0 {
		return paths
	}

	if i == m-1 && j == n-1 {
		paths = append(paths, path)
	}

	visited[i][j] = true
	path1 := ratInMaze(i+1, j, m, n, a, visited, path+"D")
	path2 := ratInMaze(i, j-1, m, n, a, visited, path+"L")
	path3 := ratInMaze(i, j+1, m, n, a, visited, path+"R")
	path4 := ratInMaze(i-1, j, m, n, a, visited, path+"T")
	visited[i][j] = false

	paths = append(paths, path1...)
	paths = append(paths, path2...)
	paths = append(paths, path3...)
	paths = append(paths, path4...)

	return paths
}
