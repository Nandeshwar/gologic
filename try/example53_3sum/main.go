// https://www.youtube.com/watch?v=fpq-Sz_tpzM
package main

import (
	"sort"
)

func ThreeSum(numArr []int) [][]int {
	result := [][]int{}

	sort.Ints(numArr)
	for i := 0; i < len(numArr)-2; i++ {
		twoSum := findTwoSum(0-numArr[i], numArr, i+1)
		if len(twoSum) == 2 {
			threeSum := []int{numArr[i], twoSum[0], twoSum[1]}
			result = append(result, threeSum)
		}
	}

	return result
}

func findTwoSum(v int, numList []int, i int) []int {
	j := len(numList) - 1
	for i < j {
		if numList[i]+numList[j] < v {
			i++
		} else if numList[i]+numList[j] > v {
			j--
		} else {
			return []int{numList[i], numList[j]}
		}
	}
	return []int{}
}
