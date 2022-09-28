package main

import (
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

var counter int

func main() {
	fmt.Println("recursive")

	pre := make([]int, 5)
	post := make([]int, 5)

	dfsRecurssion(pre, post)
	fmt.Println("pre =", pre)
	fmt.Println("post=", post)
}

func dfsR(v int, visited []bool, pre, post []int) {
	visited[v] = true

	fmt.Println(v)

	pre[v] = counter
	counter++

	if v < len(Graph) {
		for _, u := range Graph[v] {
			if !visited[u] {
				dfsR(u, visited, pre, post)
			}
		}
	}

	post[v] = counter
	counter++
}

func dfsRecurssion(pre, post []int) {
	visited := make([]bool, 5)

	for i := 0; i < 5; i++ {
		if !visited[i] {
			dfsR(i, visited, pre, post)
		}
	}
}
