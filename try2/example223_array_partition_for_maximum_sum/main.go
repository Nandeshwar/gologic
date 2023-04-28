package main

import (
	"fmt"
)

/*
Given an integer array arr, partition the array into (contiguous) subarrays of length at most k. After partitioning,
each subarray has their values changed to become the maximum value of that subarray.

Example 1:

Input: arr = [1,15,7,9,2,5,10], k = 3
Output: 84
Explanation: arr becomes [15,15,15,9,10,10,10]
Example 2:

Input: arr = [1,4,1,5,7,3,6,1,9,9,3], k = 4
Output: 83
Example 3:

Input: arr = [1], k = 1
Output: 1

*/

func main() {
	a := []int{1, 15, 7, 9, 2, 5, 10}
	k := 3

	fmt.Println("maxSum array partition=", maxSubArrayPartition(a, k, 0))
}

func maxSubArrayPartition(a []int, k, ind int) int {
	if ind == len(a) {
		return 0
	}

	l := 0
	currentMax := 0
	maxSum := 0

	// front direction partitioning
	// paritition from 0 , 1 ... < min of ind+k or len of a
	// l: for loop moves front l++.   l * currentMax ----> 15 + 15 + 15 example
	for i := ind; i < min(ind+k, len(a)); i++ {
		l++
		currentMax = max(currentMax, a[i])
		sum := l*currentMax + maxSubArrayPartition(a, k, i+1)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
