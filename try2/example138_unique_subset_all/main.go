package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 2}
	// expectation: no duplicate sub arrays
	// and in order
	// and item should not be repeated in sub-arrays
	// output:
	/*
	   [[] [1] [1 2] [1 2 2] [2] [2 2]]
	*/

	ds := []int{}
	ans := [][]int{}
	ind := 0
	sort.Ints(a)
	subset(a, ind, ds, &ans)
	fmt.Println(ans)
}

func subset(a []int, ind int, ds []int, ans *[][]int) {

	t := make([]int, len(ds))
	copy(t, ds)
	*ans = append(*ans, t)

	for i := ind; i < len(a); i++ {
		//  iteration: pick one item, and if 2nd item is same, do not pick that to avoid duplicate result
		if i > ind && a[i] == a[i-1] {
			continue
		}

		ds = append(ds, a[i])
		subset(a, i+1, ds, ans)
		ds = ds[:len(ds)-1]
	}

}
