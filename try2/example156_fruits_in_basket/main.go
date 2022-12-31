package main

import (
	"fmt"
)

func main() {
	// Array is tree:
	// different number indicates different fruit
	// fruits can be picked sequentially
	// each basket contains similar fruits only
	// how many max fruits can be collected in two baskets
	a := []int{1, 2, 3, 2, 2, 2, 1, 3} // 5
	// a = []int{1, 2, 1}       // 3
	// a = []int{0, 1, 2, 2}    // 3
	// a = []int{1, 2, 3, 2, 2} // 4
	// a = []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4} // 5
	basket := 2
	fmt.Println("original fruits=", a)
	totalFruitsIn2Baskets := fruitsInBasket(a, basket)
	fmt.Println("totalFruitsIn2Baskets=", totalFruitsIn2Baskets)
	/*
		output:
		original fruits= [1 2 3 2 2 2 1 3]
		totalFruitsIn2Baskets= 11
	*/
}

// Sliding window technique
func fruitsInBasket(a []int, basket int) int {
	var maxFruits int
	m := make(map[int]int)

	for i := 0; i < basket; i++ {
		v, ok := m[a[i]]
		if ok {
			m[a[i]] = v + 1
		} else {
			m[a[i]] = 1
		}
	}

	for i := basket; i < len(a); i++ {

		fmt.Println("m=", m)

		maxFruits = max(maxFruits, countFruits(m))

		nextItem, ok := m[a[i]]
		if ok {
			m[a[i]] = nextItem + 1
		} else {
			m[a[i]] = 1
		}

		if len(m) > 2 {

			firstItemInWindow := i - basket
			delete(m, a[firstItemInWindow])
		}
		maxFruits = max(maxFruits, countFruits(m))

	}

	return maxFruits
}

func countFruits(m map[int]int) int {
	totalFruitsInBasket := 0
	for _, v := range m {
		totalFruitsInBasket += v
	}
	return totalFruitsInBasket
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
