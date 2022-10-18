package main

import "fmt"

func main() {
	a := intersection([]int{10, 20, 20, 30, 40}, []int{20, 30})
	fmt.Println(a)
}

func intersection(a, b []int) []int {
	m := map[int]struct{}{}
	for _, v := range a {
		m[v] = struct{}{}
	}

	intersectionMap := map[int]struct{}{}
	for _, v := range b {
		_, ok := m[v]
		if ok {
			intersectionMap[v] = struct{}{}
		}
	}

	resultArr := make([]int, len(intersectionMap))
	ind := 0
	for k, _ := range intersectionMap {
		resultArr[ind] = k
		ind++
	}

	return resultArr
}
