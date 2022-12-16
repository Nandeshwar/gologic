package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 1}
	target := 2
	ind := 0
	sum := 0
	/* target sub arrays
		output:
		 [1 1]
		 [2]

	  all sub arrays
	  --------------
	[1 1]
	[2]
	[1 2 1]
	*/

	tmp := []int{}
	fmt.Println("All sub arrays")
	allSubArrays(a, tmp, ind)
	fmt.Println("target sub arrays")
	targetSubArraysAll(a, target, tmp, ind, sum)
	fmt.Println("any 1 sub-arrays")
	targetSubArraysOne(a, target, tmp, ind, sum, false)
	fmt.Println("Any 1 algo 2")
	targetSubArraysOneAlgo2(a, target, tmp, ind, sum)
	fmt.Println("target subArrays count")
	fmt.Println(targetSubArraysCount(a, target, ind, sum))
}

func allSubArrays(a []int, tmp []int, ind int) {
	if ind == len(a) {
		fmt.Println(tmp)
		return
	}

	
	tmp = append(tmp, a[ind])
	allSubArrays(a, tmp, ind+1)

	tmp = tmp[0 : len(tmp)-1]

	allSubArrays(a, tmp, ind+1)
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

func targetSubArraysCount(a []int, target int, ind int, sum int) int {
	if ind == len(a) {
		if target == sum {
			return 1
		}
		return 0
	}

	sum += a[ind]
	left := targetSubArraysCount(a, target, ind+1, sum)

	sum -= a[ind]

	right := targetSubArraysCount(a, target, ind+1, sum)

	return left + right

}
