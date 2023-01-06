package main

import (
	"fmt"
)

func main() {
	coins := []int{1, 3, 2}
	target := 3

	var result [][]int
	allCombinationArr(coins, target, []int{}, &result)
	fmt.Println("result=", result)
	/*
			output:
		[1 1 1]
		[1 2]
		[3]
		[2 1]
	*/
}

func allCombinationArr(coins []int, target int, ds []int, result *[][]int) {

	if target == 0 {
		fmt.Println(ds)
		tmp := make([]int, len(ds))
		copy(tmp, ds)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(coins); i++ {
		if target-coins[i] >= 0 {
			ds = append(ds, coins[i])
			allCombinationArr(coins, target-coins[i], ds, result)
			ds = ds[0 : len(ds)-1]
		}
	}

}