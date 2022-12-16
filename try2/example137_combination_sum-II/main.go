package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 1, 1, 5, 3, 2}
	target := 5
	// expectation: no duplicate sub arrays
	// and in order
	// and item should not be repeated in sub-arrays
	// output:
	/*
	   [[1 1 1 2] [1 1 3] [2 3] [5]]
	*/

	ds := []int{}
	ans := [][]int{}
	ind := 0
	sort.Ints(a)
	combinationSum(a, ind, target, ds, &ans)
	fmt.Println(ans)
}

func combinationSum(a []int, ind int, target int, ds []int, ans *[][]int) {

	if target == 0 {
		t := make([]int, len(ds))
		copy(t, ds)
		*ans = append(*ans, t)
		return
	}

	for i := ind; i < len(a); i++ {
		if i > ind && a[i] == a[i-1] {
			continue
		}
		if a[i] > target {
			break
		}
		ds = append(ds, a[i])
		combinationSum(a, i+1, target-a[i], ds, ans)
		ds = ds[:len(ds)-1]
	}

}
