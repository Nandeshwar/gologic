package main

import (
	"fmt"
)

func main() {
	a := []int{6, 4, 3, 1, 2}
	evenFirstOddNext(a)
	fmt.Println(a)

	/*
		output:
		  [6 4 2 1 3]
	*/
}

func evenFirstOddNext(a []int) {
	i := 0
	j := len(a) - 1

	for i < j {
		if a[i]%2 == 0 {
			i++
			continue
		}

		if a[j]%2 == 1 {
			j--
			continue
		}

		if i < j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}

	}
}
