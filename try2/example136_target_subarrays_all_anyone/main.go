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
	fmt.Println("any 1 sub-arrays")
	targetSubArraysOne(a, target, tmp, ind, sum, false)
	fmt.Println("Any 1 algo 2")
	targetSubArraysOneAlgo2(a, target, tmp, ind, sum)
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
func targetSubArraysOneAlgo2(a []int, target int, tmp []int, ind int, sum int) bool {
	if ind == len(a) {
		if target == sum {
			fmt.Println(tmp)
			return true
		}
		return false
	}

	sum += a[ind]
	tmp = append(tmp, a[ind])
	left := targetSubArraysOneAlgo2(a, target, tmp, ind+1, sum)
	if left {
		return true
	}
	sum -= a[ind]

	tmp = tmp[0 : len(tmp)-1]

	right := targetSubArraysOneAlgo2(a, target, tmp, ind+1, sum)
	if right {
		return true
	}
	return false

}
func targetSubArraysOne(a []int, target int, tmp []int, ind int, sum int, found bool) bool {
	if ind == len(a) {
		if target == sum {
			fmt.Println(tmp)
			found = true
		}
		return found
	}

	if found {
		return found
	}

	sum += a[ind]
	tmp = append(tmp, a[ind])
	found = targetSubArraysOne(a, target, tmp, ind+1, sum, found)
	sum -= a[ind]

	tmp = tmp[0 : len(tmp)-1]

	found = targetSubArraysOne(a, target, tmp, ind+1, sum, found)
	return found

}
