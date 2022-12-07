package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 5}
	target := 8

	result := [][]int{}
	ds := []int{}
	ind := 0
	getAllArraysToTarget(a, ind, target, &result, &ds)
	fmt.Println("result=", result)
}

func getAllArraysToTarget(a []int, ind int, target int, result *[][]int, ds *[]int) {

	if ind == len(a) {

		if target == 0 {
			fmt.Println("(*ds)=", *ds)
			d := make([]int, len(*ds))
			copy(d, *ds)
			*result = append(*result, d)
		}
		return
	}
	if a[ind] <= target {
		*ds = append(*ds, a[ind])

		getAllArraysToTarget(a, ind, target-a[ind], result, ds)
		*ds = (*ds)[:len(*ds)-1]
	}

	getAllArraysToTarget(a, ind+1, target, result, ds)

}

/*
(*ds)= [1 1 1 1 1 1 1 1]
(*ds)= [1 1 1 1 1 3]
(*ds)= [1 1 1 5]
(*ds)= [1 1 3 3]
(*ds)= [3 5]
result= [[1 1 1 1 1 1 1 1] [1 1 1 1 1 3] [1 1 1 5] [1 1 3 3] [3 5]]
*/
