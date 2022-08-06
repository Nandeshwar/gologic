package main

func findMaxSubArr(arr []int) int {
	maxSum := 0
	currSum := 0

	for i, _ := range arr {
		currSum += arr[i]

		if currSum > maxSum {
			maxSum = currSum
		}

		if currSum < 0 {
			currSum = 0
		}
	}
	return maxSum
}
