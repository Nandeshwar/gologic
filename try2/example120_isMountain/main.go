package main

import (
	"fmt"
)

func main() {
	a := []int{3, 4, 5, 3}
	fmt.Println(isMountain(a))
}

func isMountain(a []int) bool {
	if len(a) < 3 {
		return false
	}

	i := 0
	// false when it reaches to end i.e number is increasing
	for i < len(a) && i+1 < len(a) {
		if a[i] < a[i+1] {
			i++
		} else {
			break
		}
	}

	fmt.Println(i)

	if i == len(a)-1 {
		return false
	}

	// false when last two numbers are same
	if a[i] == a[i+1] {
		return false
	}

	return true
}
