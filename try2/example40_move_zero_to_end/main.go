package main

import "fmt"

func main() {
	a := []int{0, 1, 0, 3, 2}

	index := 0
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			a[index] = a[i]
			index++
		}
	}

	for i := index; i < len(a); i++ {
		a[i] = 0
	}

	fmt.Println(a)
}
