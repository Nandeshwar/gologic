package main

import "fmt"

func main() {
	k := 3
	target := 9

	// Find all unique sub arrays of length 3 of sum num from array between 1 to 9
	findAllUniqueSubArrays(k, 1, target, []int{})

	/*
				output:
		[1 2 6]
		[1 3 5]
		[2 3 4]
	*/
}

func findAllUniqueSubArrays(k, ind, target int, ds []int) {

	if len(ds) == k {
		if target == 0 {
			fmt.Println(ds)
		}
		return
	}

	for i := ind; i <= 9; i++ {
		if i > target {
			break
		}

		ds = append(ds, i)
		findAllUniqueSubArrays(k, i+1, target-i, ds)
		ds = ds[:len(ds)-1]
	}
}
