package main

import (
	"fmt"
)

func main() {

	fmt.Println("total path=", countPath(3, 3))
}

func countPath(r, c int) int {
	if r == 1 || c == 1 {
		return 1
	}

	return countPath(r-1, c) + countPath(r, c-1)
}
