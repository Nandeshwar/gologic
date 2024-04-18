package main

import "fmt"

func main() {
	fmt.Println(algo1([]int{100, 200, 4, 2, 1, 3}))
}

func algo1(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	maxCnt := 0
	for _, v := range nums {
		cnt := 1
		nextVal := v

		// if the previous value exist, that means it's not starting point and continue
		// example: 100 200 2 4 1 3
		// 100 no previous value this starting point
		// 200 no previous value starting point
		// 2 has previous value in map so not starting point so proceed with next item
		// 4 has previous value proceed with next item
		// 1 does not have previous value. so run logic given below
		_, ok := m[v-1]
		if ok {
			continue
		}

		// increase count if next value exists
		for {
			nextVal++
			_, ok := m[nextVal]
			if ok {
				cnt++
				maxCnt = Max(maxCnt, cnt)
			} else {
				maxCnt = Max(maxCnt, cnt)
				break
			}

		}
	}

	return maxCnt
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
