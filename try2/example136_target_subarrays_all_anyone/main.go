package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 1}
	target := 2
	ind := 0
	sum := 0
	/*
		output:
		 [1 1]
		 [2]
	*/

	tmp := []int{}
	targetSubArraysAll(a, target, tmp, ind, sum)
}

func targetSubArraysAll(a []int, target int, tmp []int, ind int, sum int) {
	if ind == len(a) {
		if target == sum {
			fmt.Println(tmp)
		}
		return
	}

	sum += a[ind]
	tmp = append(tmp, a[ind])
	targetSubArraysAll(a, target, tmp, ind+1, sum)
	sum -= a[ind]

	tmp = tmp[0 : len(tmp)-1]

	targetSubArraysAll(a, target, tmp, ind+1, sum)
}
