package main

import (
	"fmt"
	"sort"
)

/*
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
*/

func main() {
	a := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(a))

}

func threeSum(a []int) [][]int {
	sort.Ints(a)
	fmt.Println("a=", a)
	m := make(map[[3]int]struct{})

	for i := 0; i < len(a)-2; i++ {
		j := i + 1
		k := len(a) - 1
		for j < k {
			if a[i]+a[j]+a[k] == 0 {
				result := [3]int{a[i], a[j], a[k]}

				sort.Ints(result[:])
				m[result] = struct{}{}
			}
			j++
			k--
		}
	}

	var ans [][]int
	for k, _ := range m {
		kk := []int{k[0], k[1], k[2]}
		ans = append(ans, kk)
	}
	return ans
}
