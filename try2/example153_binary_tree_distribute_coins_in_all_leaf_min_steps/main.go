package main

import (
	"fmt"
	"math"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func main() {
	/*
			3
		  0    0

		expectation: 2
	*/
	zero2 := &Tree{item: 0}
	zero1 := &Tree{item: 0}
	t := &Tree{item: 3, left: zero1, right: zero2}
	ans := 0
	distributeCoinsMinSteps(t, &ans)
	fmt.Println("minSteps=", ans)
}

func distributeCoinsMinSteps(t *Tree, ans *int) int {
	if t == nil {
		return 0
	}
	left := distributeCoinsMinSteps(t.left, ans)
	right := distributeCoinsMinSteps(t.right, ans)

	*ans += int(math.Abs(float64(left)) + math.Abs(float64(right)))
	return t.item - 1 + left + right
}
