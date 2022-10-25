package main

import (
	"container/list"
	"fmt"
)

func main() {
	keys := [][]int{
		{1, 3},
		{},
		{},
		{2},
	}

	fmt.Println(roomVisited(keys))
}

func roomVisited(keys [][]int) bool {
	visited := make([]bool, len(keys))

	visited[0] = true
	stack := list.New()
	stack.PushBack(0)

	for stack.Len() != 0 {
		element := stack.Remove(stack.Back())
		index := element.(int)

		for _, key := range keys[index] {
			if !visited[key] {
				stack.PushBack(key)
				visited[key] = true
			}
		}
	}

	for _, v := range visited {
		if v == false {
			return false
		}
	}

	return true
}
