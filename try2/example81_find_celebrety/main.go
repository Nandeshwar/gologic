package main

import (
	"fmt"
)

func main() {

	// everyone know celebrety. Celebrety does not know anyone
	fmt.Println(findCelebrety(5))
}

func findCelebrety(n int) int {
	candidate := 0
	for i := 1; i < n; i++ {
		if know(candidate, i) {
			fmt.Println("know....", candidate, " ", i)
			candidate = i // i holds next item and that will new candidate
		}
	}

	fmt.Println("candidate=", candidate)

	for i := 0; i < n; i++ {
		if i != candidate && (know(candidate, i) || !know(i, candidate)) {
			return -1
		}
	}
	return candidate
}

func know(candidate, i int) bool {
	if candidate == 3 {
		return false
	}
	return true
}

/*
output
bash-3.2$ go run main.go
know.... 0   1
know.... 1   2
know.... 2   3
candidate= 3
3
*/
