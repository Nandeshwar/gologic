package main

import (
	"fmt"
)

func main() {
	fmt.Println(stock2BuySell([]int{4, 2, 2, 2, 4, 8, 9}))
	fmt.Println(stock2BuySellAlgo2([]int{4, 2, 2, 2, 4, 8, 9}))
}

// buy/sell stock multiple times. But can not have more than 1 stock.
// check price if greater than left sell
func stock2BuySell(a []int) int {
	if len(a) == 1 {
		return 0
	}
	totalProfit := 0
	for i := 1; i < len(a); i++ {
		if a[i-1] < a[i] {
			currProfit := a[i] - a[i-1]
			totalProfit += currProfit
		}
	}
	return totalProfit
}

// check price if greater than left - mark it minimum, keep checking further for highest price
func stock2BuySellAlgo2(a []int) int {
	if len(a) == 1 {
		return 0
	}
	totalProfit := 0
	var min int
	var lock bool
	for i := 1; i < len(a); i++ {
		if a[i-1] < a[i] {
			if lock == false {
				min = a[i-1]
				lock = true
			}

			if i+1 == len(a) || a[i] > a[i+1] {
				currentProfit := a[i] - min
				totalProfit += currentProfit
				min = 0
				lock = false
			}
		}
	}
	return totalProfit
}
