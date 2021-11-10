/*
input:
   n = 5 : only take 5 items
   v = 15 14 10 45 30
   w = 2   5 1   3 4
   bag weight = 7

   max weight of item will be same. choose weight which value is max

  understand the problem from this link: https://www.youtube.com/watch?v=bUSaenttI24

  and nice explanation of code using recursion: https://www.youtube.com/watch?v=kvyShbFVaY8

*/
package main

import (
	"github.com/logic-building/functional-go/fp"
)

func knapsackProblem(w, v []int, bagWeight, n int) int {
	// base condition
	if n == 0 || bagWeight == 0 {
		return 0
	}

	// if weight < bagweight
	// we have 2 option a. include value or not include. so take max of both
	if w[n-1] <= bagWeight {

		includedVal := v[n-1] + knapsackProblem(w, v, bagWeight-w[n-1], n-1) // include and proceed for other items
		notIncludedVal := knapsackProblem(w, v, bagWeight, n-1) // not include and proceed for other items

		return fp.MaxInt([]int{includedVal, notIncludedVal})
	} else { // if weight > bagWeight then not include item and proceed of rest of the items
		return knapsackProblem(w, v, bagWeight, n-1)
	}

}
