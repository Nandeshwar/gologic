package main

import "fmt"

func main() {
	fmt.Println(max([]int{1, 20, 10, 5, 9}, 0))
}

func max(a []int, ind int) int {
	if ind == len(a)-1 {
		return a[ind]
	}

	m := max(a, ind+1)
	if m > a[ind] {
		return m
	}
	return a[ind]
}
