package main

import "fmt"

func main() {
	fmt.Println(max([]int{1, 20, 10, 5, 9}, 0))
	fmt.Println(max2([]int{1, 20, 10, 5, 9}, 0))
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

func max2(a []int, ind int) int {
	if len(a) == 1 {
		return a[0]
	}

	m := max2(a[1:], ind+1)
	if m > a[0] {
		return m
	}
	return a[0]
}
