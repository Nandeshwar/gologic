package main

import "fmt"

func main() {
	fmt.Println("hello")

	a := []int{10, 5, 2, 3, 3, 3}

	m := map[int]struct{}{}

	for _, v := range a {
		_, ok := m[v]
		if ok {
			fmt.Println("duplicate found ")
			break;
		}
		m[v] = struct{}{}
	}
}