package main

import (
	"container/list"
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

func main() {
	displayGraph(0)
}

func displayGraph(v int) {
	visited := []bool{false, false, false, false, false}
	l := list.New()
	l.PushBack(v)
	visited[v] = true

	for l.Len() != 0 {
		element := l.Remove(l.Front())
		v := element.(int)
		fmt.Println(v)

		if v >= len(Graph) {
			continue
		}

		for _, item := range Graph[v] {
			if visited[item] == false {
				l.PushBack(item)
				visited[item] = true
			}
		}
	}

}
