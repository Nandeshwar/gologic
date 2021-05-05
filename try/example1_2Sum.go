package main

import "fmt"

// python example  explanation - https://www.youtube.com/watch?v=V1Yqw_ge1IM
func main() {
	fmt.Println(twoSum())
	fmt.Println(twoSum2())

}

// Find index of numbers if their sum is equal to target(here: 9)
// time: o(n2)
// space: o(1)
func twoSum() (int, int) {
	target := 9
	arr1 := []int{2, 7, 9, 11}

	for i, v := range arr1 {
		for j := i + 1; j <= len(arr1); j++ {
			if v+arr1[j] == target {
				return i, j
			}
		}
	}
	return 0, 0
}


// Find index of numbers if their sum is equal to target(here: 9)
// time: o(n)
// space: o(n)
func twoSum2() (int, int) {
	target := 9
	arr1 := []int{2, 7, 9, 11}

	m := map[int]int{}

	for i, v := range arr1 {
		indexOfItemInMap, ok := m[v]
		if ok {
			return indexOfItemInMap, i
		}
		m[target-v] = i
		
	}
	return 0, 0
}