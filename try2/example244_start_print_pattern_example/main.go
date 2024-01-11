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

	fmt.Println("------------")
	pattern2()
}

func pattern2() {
	/*
		Hello World
		E
		DE
		CDE
		BCDE
		ABCDE
	*/
	fmt.Println("Hello World")
	n := 5

	for i := 0; i < n; i++ {
		var b int32
		b = int32(n - i - 1)
		a := 'A'
		a = a + b
		for j := 0; j <= i; j++ {
			fmt.Print(string(a))
			a++
		}
		fmt.Println()
	}

}
