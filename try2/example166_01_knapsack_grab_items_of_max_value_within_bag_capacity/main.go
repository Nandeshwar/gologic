package main

import (
	"fmt"
)

func main() {
	weights := []int{3, 5, 8, 4, 4}
	values := []int{20, 10, 5, 40, 50}
	bagCapacity := 8

	// expectation: 90

	m := make(map[string]int)
	maxValues := findMaxValues(weights, values, bagCapacity, len(weights)-1, m)
	fmt.Println("max values=", maxValues)

	maxValues2 := findMaxValues2(weights, values, bagCapacity, len(weights)-1, m)
	fmt.Println("max values=", maxValues2)
}

func findMaxValues(weights, values []int, bagCapacity, ind int, m map[string]int) int {
	if ind == 0 {
		if weights[ind] <= bagCapacity {
			return values[0]
		} else {
			return 0
		}
	}

	v, ok := m[string(ind)+string(bagCapacity)]
	if ok {
		return v
	}
	notPick := 0 + findMaxValues(weights, values, bagCapacity, ind-1, m)
	pick := -1
	if weights[ind] <= bagCapacity {
		pick = values[ind] + findMaxValues(weights, values, bagCapacity-weights[ind], ind-1, m)
	}

	maxValue := max(pick, notPick)
	m[string(ind)+string(bagCapacity)] = maxValue
	return maxValue
}

func findMaxValues2(weights, values []int, bagCapacity, ind int, m map[string]int) int {
	if ind == 0 {
		if weights[ind] <= bagCapacity {
			return values[0]
		} else {
			return 0
		}
	}
	if ind < 0 {
		return 0
	}

	v, ok := m[string(ind)+string(bagCapacity)]
	if ok {
		return v
	}
	var pick int
	var notPick int
	if weights[ind] <= bagCapacity {
		notPick = values[ind] + findMaxValues2(weights, values, bagCapacity-weights[ind], ind-1, m)
		pick = values[ind-1] + findMaxValues2(weights, values, bagCapacity-weights[ind], ind-2, m)
	}

	maxValue := max(pick, notPick)
	m[string(ind)+string(bagCapacity)] = maxValue
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
