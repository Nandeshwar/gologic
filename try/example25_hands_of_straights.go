// Take 3 consecutive numbers as internal array
package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{1, 2, 3, 6, 2, 3, 4, 5, 7}
	output := make([][]int, 3)
	sort.Ints(input)
	fmt.Println(input)

	outputPointer := -1
	
	cnt := 1
	for i, v := range input {
		if v == 0 {
			continue
		}
		if cnt % 3 == 0 || cnt == 1 {
			outputPointer++
			output[outputPointer] = make([]int, 3)
		}

		valueOfI := input[i]
		
		for k := 0; k < 3; k++ {
			for l := i; l < len(input); l++ {
				if valueOfI + k  == input[l] {
					output[outputPointer][k] = input[l]
					input[l] = 0
					break
				}
			}
		}
	}

	fmt.Println(output)
}
/*
output:
  [1 2 2 3 3 4 5 6 7]
[[1 2 3] [2 3 4] [5 6 7]]
*/