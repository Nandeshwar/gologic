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
			candidate = i
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
