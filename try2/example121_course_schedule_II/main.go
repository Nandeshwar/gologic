package main

import (
	"container/list"
	"fmt"
)

var Graph [][]int

/*
 0 ------>1
 |       |
 |       |
 |       |
 D       D
 2-----> 3
*/

var courseSchedule = [][]int{
	{1, 0}, // 0 course then 1
	{2, 0}, // 0 course then 2
	{3, 1},
	{3, 2},
}

// This has cycle
// var courseSchedule = [][]int{
// 	{0, 1},
// 	{1, 0},
// }

type Status int

const (
	NotVisited Status = iota
	Visited
	InStack
)

func main() {
	stack := list.New()
	status := make([]Status, len(courseSchedule))
	createGraph()
	fmt.Println(Graph)
	list := findCourseSchedule(stack, status)
	fmt.Println("Result=", list)

}

func findCourseSchedule(stack *list.List, status []Status) []int {
	var result []int

	for i := 0; i < 4; i++ {

		if status[i] == NotVisited && hasCycle(i, stack, status) {
			return []int{}
		}
	}

	for stack.Len() != 0 {
		element := stack.Remove(stack.Back())
		item := element.(int)
		result = append(result, item)
	}

	return result
}

func hasCycle(v int, stack *list.List, status []Status) bool {
	status[v] = InStack

	for _, u := range Graph[v] {
		if status[u] == InStack {
			return true
		}
		if status[u] == NotVisited && hasCycle(u, stack, status) {
			return true
		}
	}

	status[v] = Visited
	stack.PushBack(v)
	return false
}

func createGraph() {
	Graph = make([][]int, len(courseSchedule))

	for _, s := range courseSchedule {
		Graph[s[1]] = append(Graph[s[1]], s[0])
	}
}

/*
output:
bash-3.2$ go run main.go
[[1 2] [3] [3] []]
Result= [0 2 1 3]
*/
