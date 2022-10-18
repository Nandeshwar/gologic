package main

import (
	"fmt"
)

func main() { 
   // 1 is blocker
	a := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{0, 0, 0},
	}
	// output:path1:  bbrr (bottom-> bottom -> right -> right
    //        path2:  brbr ( bottom -> right -> bottom -> right)

	printAllPaths(a)
}

func printAllPaths(a [][]int) {

	visited := make([][]bool, len(a))
	for i, _ := range visited {
		visited[i] = make([]bool, len(a))
	}

	tlbr(a, 0, 0, visited, "")

}

func tlbr(a [][]int, r, c int, visited [][]bool, result string) {
	l := len(a)

	if r < 0 || r == l || c < 0 || c == l || a[r][c] == 1 || visited[r][c] {
		return
	}

	if r == l-1 && c == l-1 {
		fmt.Println(result)
		return
	}

	visited[r][c] = true
	tlbr(a, r-1, c, visited, result+"t")
	tlbr(a, r, c-1, visited, result+"l")
	tlbr(a, r+1, c, visited, result+"b")
	tlbr(a, r, c+1, visited, result+"r")
	visited[r][c] = false

}
