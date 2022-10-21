package main

import (
	"fmt"
)

func main() {
	subArraysOfTargetSum([]int{10, 20, 30, 40, 50}, 50)
}

func subArraysOfTargetSum(a []int, targetSum int) {
	subArrays(a, 0, targetSum, []int{}, 0)
}

func subArrays(a []int, n, targetSum int, result []int, sum int) {

	if n == len(a) {
		if sum == targetSum {
			fmt.Println(result)
		}
		return
	}

	newResult := append(result, a[n])
	newSum := sum + a[n]
	subArrays(a, n+1, targetSum, newResult, newSum)
	subArrays(a, n+1, targetSum, result, sum)
}
