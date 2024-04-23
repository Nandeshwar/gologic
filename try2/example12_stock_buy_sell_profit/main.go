package main

import "fmt"

func main() {
	/*
	   	Example 1:

	      Input: prices = [7,1,5,3,6,4]
	      Output: 5

	*/

	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
func maxProfit(prices []int) int {

	m := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		min = Min(min, prices[i])

		benefits := prices[i] - min
		m = Max(m, benefits)
	}
	return m

	/* Running example two pointers
	   l := 0
	   r := 1

	   m := 0
	   for r < len(prices) {
	       m = Max(m, prices[r]-prices[l])
	       if prices[l] > prices[r] {
	           l++
	       } else {
	            r++
	       }

	   }
	   return m
	*/
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
