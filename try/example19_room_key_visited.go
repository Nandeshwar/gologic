// https://www.youtube.com/watch?v=Rz_-Kx0LN-E&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=28

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// roomAndKeys := [][]int{
	// 	{1},
	// 	{2},
	// 	{3},
	// 	{},
	// }

	roomAndKeys := [][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	}


	visited := make([]bool, len(roomAndKeys))

	stack := list.New()
	stack.PushBack(0)
	visited[0] = true

	for stack.Len() > 0 {
		item := stack.Remove(stack.Back())

		for _, key := range roomAndKeys[item.(int)] {
			if !visited[key] {
				visited[key] = true
				stack.PushBack(key)
			}
		}
	}

	isVisited := true
	for _, v := range visited {
		if !v {
			isVisited = false
		}
	}
	if isVisited {
		fmt.Println("All room visited")
		return
	}
	fmt.Println("All room is not visited")

}
