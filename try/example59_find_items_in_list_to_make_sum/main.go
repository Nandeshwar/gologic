// https://www.youtube.com/watch?v=S3rnLLHl0PM&list=PLKKfKV1b9e8pWy_UIiJlOlX_T4al_UtQJ&index=5
package main

import (
	"github.com/logic-building/functional-go/fp"
)

// partition array so that their sum is equal to sum of big array
func partitionArraySum(list []int) [][]int {
	s := func(a, b int) int {
		return a + b
	}

	sum := fp.ReduceInt(s, list)

	if sum%2 != 0 {
		return [][]int{}
	}
	resultList := [][]int{}

	for len(list) > 0 {

		l := findList(list, 0, sum/2, []int{})
		if len(l) == 0 {
			break
		}
		list = fp.DifferenceInt(list, l)
		resultList = append(resultList, l)
	}
	return resultList
}

func findList(list []int, i int, sum int, result []int) []int {
	if (i >= len(list) && sum != 0) || sum < 0 {
		return []int{}
	}

	if sum == 0 {
		return result
	}
	result = append(result, list[i])

	l := findList(list, i+1, sum-list[i], result)
	if len(l) > 0 {
		return l
	}

	result = fp.DropInt(list[i], result)

	l = findList(list, i+1, sum, result)
	return l
}
