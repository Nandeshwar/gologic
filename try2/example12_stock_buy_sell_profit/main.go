package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(stockBuySellProfitAlgo1([]int{5, 1, 8})) // expectation 7
	fmt.Println(stockBuySellProfitAlgo2([]int{5, 1, 8}))
}

// each value in array represents stock buy and sell
// find profit
// preprocessing: find max value from right side
// then find max benefit
// o(n) but uses spaces
func stockBuySellProfitAlgo1(arr []int) int {
	maxProfit := 0

	maxFromRight := make([]int, len(arr)) // for preprocessing
	max := arr[len(arr)-1]

	for i := len(arr) - 1; i >= 0; i-- {
		max = int(math.Max(float64(max), float64(arr[i])))
		maxFromRight[i] = max

		profit := maxFromRight[i] - arr[i]
		maxProfit = int(math.Max(float64(maxProfit), float64(profit)))
	}
	return maxProfit
}

// o(n)
// find minimum for each index
// profit = currVal - min
func stockBuySellProfitAlgo2(arr []int) int {
	maxProfit := 0
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		min = int(math.Min(float64(min), float64(arr[i])))

		profit := arr[i] - min
		maxProfit = int(math.Max(float64(maxProfit), float64(profit)))
	}
	return maxProfit
}
