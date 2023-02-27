package main

import (
	"fmt"
)

func main() {
	a := []int{1, 7, 5} // how many coins to make it 18
	coinCnt := coinCount(a, 18)
	fmt.Println("total coin count=", coinCnt)
	fmt.Println("using dynamic programming")
	coinCnt = coinCountDp(a, 18)
	fmt.Println("total coin count=", coinCnt)

	fmt.Println("Algorithm2....")
	fmt.Println(coinCount2(a, 18, 0, []int{}))
}

func coinCount(a []int, sum int) int {
	ans := int(^uint(0) >> 1)
	if sum == 0 {
		return 0
	}

	for _, v := range a {
		if sum-v >= 0 {
			subAns := coinCount(a, sum-v) + 1
			ans = min(ans, subAns)
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var m = map[int]int{}

func coinCountDp(a []int, sum int) int {
	ans := int(^uint(0) >> 1)
	if sum == 0 {
		return 0
	}

	for _, v := range a {
		if sum-v >= 0 {
			subAns := 0
			sv, ok := m[sum]
			if ok {
				return sv
			}
			subAns = coinCountDp(a, sum-v) + 1
			ans = min(ans, subAns)
		}
	}

	m[sum] = ans
	return ans
}

func coinCount2(a []int, target, ind int, ds []int) int {
	if ind == len(a) {
		ans := int(^uint(0) >> 1)
		if target == 0 {

			ans = len(ds)
		}
		return ans
	}

	ans1 := int(^uint(0) >> 1)
	if a[ind] <= target {
		ds = append(ds, a[ind])
		ans1 = coinCount2(a, target-a[ind], ind, ds)
		ds = ds[0 : len(ds)-1]
	}
	ans2 := coinCount2(a, target, ind+1, ds)

	return min(ans1, ans2)
}
