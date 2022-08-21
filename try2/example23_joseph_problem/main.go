package main

import (
	"fmt"
)

func main() {
	// 6 people, kill every 3rd one. expected answer 3
	fmt.Println(joseph(5, 3))
}

func joseph(n, k int) int {
	if n == 1 {
		return 0
	}

	x := joseph(n-1, k)
	y := (x + k) % n

	// or the above two line can be written as
	// return (joseph(n-1, k) + k) % n
	return y
}
