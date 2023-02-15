package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// a := []int{10, 2} // expectation= 210
	a := []int{3, 30, 34, 5, 9} // expectation= 9534330
	fmt.Println("largest number: ", largestNumber(a))
}

func largestNumber(a []int) string {
	sort.Slice(a, func(i, j int) bool {
		return strconv.Itoa(a[i])+strconv.Itoa(a[j]) > strconv.Itoa(a[j])+strconv.Itoa(a[i])
	})

	var result string
	for _, v := range a {
		result += strconv.Itoa(v)
	}
	return result
}
