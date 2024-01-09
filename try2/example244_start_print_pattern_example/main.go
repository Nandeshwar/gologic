package main

import (
	"fmt"
)

func main() {
	/*
		 	*
		   ***
		  *****
		 *******
		*********
	*/

	// outer for loop
	for i := 1; i <= 5; i++ {

		// print space - 4, 3, 2, 1
		for j := 0; j < 5-i; j++ {
			fmt.Print(" ")
		}

		// print star:
		for j := 0; j < i+i-1; j++ {
			fmt.Print("*")
		}

		// print space - 4, 3, 2, 1
		for j := 0; j < 5-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
