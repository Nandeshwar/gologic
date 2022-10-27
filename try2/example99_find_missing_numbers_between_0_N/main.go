package main

import (
	"fmt"
)

func main() {
	a := []int{1, 0, 3, 5, 4}
	fmt.Println(findMissingNumber(a))
	fmt.Println(findMissingNumber2(a))
}

func findMissingNumber2(arr []int) int {
	//natural number always starts from 1
	// formula to find sum of natural numbers: n/2[n*a + (n-1) *d] where n total length. a is first term, d is difference between two consecutive term

	n := float64(len(arr)) //  n should be between 1 to n(here 1 to 5) but it starts from 0 and a number is missing so n=5
	a := 1.0
	d := 1.0

	sum := n / 2 * (2*a + (n-1)*d)
	// or
	sum = n * (n + 1) / 2
	fmt.Println("total natural sum=", sum)

	arrSum := 0
	for _, v := range arr {
		arrSum += v
	}

	return int(sum) - arrSum
}

func findMissingNumber(a []int) int {
	m := map[int]struct{}{}
	for _, v := range a {
		m[v] = struct{}{}
	}

	n := len(a)
	for i := 0; i <= n; i++ {
		_, ok := m[i]
		if !ok {
			return i
		}
	}
	return -1
}
