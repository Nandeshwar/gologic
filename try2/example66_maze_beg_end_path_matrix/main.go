package main

import "fmt"

func main() {
	fmt.Println(mazePath(1, 1, 3, 3))
}

func mazePath(r, c, rl, cl int) []string {
	if r == rl && c == cl {
		return []string{""}
	}

	var cPaths []string
	var rPaths []string

	if r < rl {
		rPaths = mazePath(r+1, c, rl, cl)
	}

	if c < cl {
		cPaths = mazePath(r, c+1, rl, cl)
	}

	totalPaths := []string{}
	for _, p := range cPaths {
		totalPaths = append(totalPaths, "c"+p)
	}

	for _, p := range rPaths {
		totalPaths = append(totalPaths, "r"+p)
	}

	return totalPaths
}
