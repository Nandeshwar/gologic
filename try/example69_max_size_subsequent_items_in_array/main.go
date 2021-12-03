// Kaden's Algorithm
package main

import (
	"fmt"
)

func findMaxValueAndContinousSubarry(arr []int) (max int, maxArr []int) {
	fmt.Println(arr)
	currSum := arr[0]
	maxSum := arr[0]

	currArr := []int{arr[0]}
	maxArr = append(maxArr, arr[0])

	for _, v := range arr[1:] {
		if currSum >= 0 {
			currSum += v
			currArr = append(currArr, v)
		} else {
			currSum = v
			currArr = []int{v}
		}

		if currSum > maxSum {
			maxSum = currSum
			maxArr = make([]int, len(currArr))
			copy(maxArr, currArr)
		}

	}

	return maxSum, maxArr
}
