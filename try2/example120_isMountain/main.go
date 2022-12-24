package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 4}
	fmt.Println(isMountain(a))
}

func isMountain(a []int) bool {
	if len(a) < 3 {
		return false
	}

	i := 0

	for i < len(a)-1 {
		if a[i] < a[i+1] {
			i++
		} else {
			break
		}
	}

	fmt.Println(i)
	// false when i reaches to end i.e number is increasing there no decreasing value. so no mountain
	if i == len(a)-1 {
		return false
	}

	// false when last two numbers are same
	if i+1 < len(a) && a[i] == a[i+1] {
		return false
	}

	return true
}
