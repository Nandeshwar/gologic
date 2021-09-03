/*
Sort array: even number should be in even index and odd number should be odd index

*/
package main

import (
	"fmt"
)

func main() {
	input := []int{4, 2, 5, 7}
	// output: [4 5 2 7]

	j := 1
	for i := 0; i < len(input); i += 2 {
		if input[i] % 2 == 1 {
			for j % 2 == 1 {
				input[i], input[j] = input[j], input[i]
				j++
			}
		}
	}

	fmt.Println(input)

}