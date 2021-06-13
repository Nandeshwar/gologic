// https://www.youtube.com/watch?v=MfXxic8IhkI&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=42
package main

import (
	"fmt"
	"math"
)

type Node28 struct {
	item int
	left *Node28
	right *Node28
}

func main() {
	/*
			1
		0       0
	3

	output: 4
	*/

	root := &Node28 {
		item: 1,
		left: &Node28{
			item: 0,
			left: &Node28{item: 3},
		},
		right: &Node28{item: 0},
	}

	ans := 0
	distributeCoin(root, &ans)
	fmt.Println(ans)
}

func distributeCoin(root *Node28, ans *int) int {
	if root == nil {
		return 0
	}
	left := distributeCoin(root.left, ans)
	right := distributeCoin(root.right, ans)

	*ans += int(math.Abs(float64(left)) + math.Abs(float64(right)))
	return left + right + root.item - 1
}