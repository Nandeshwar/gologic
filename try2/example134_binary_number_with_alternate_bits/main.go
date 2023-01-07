package main

import (
	"fmt"
)

func main() {
	// 5 - 101 true : there is alternate bit
	// 7 - 111 false
	// 10 - 1010 : true

	num := 10
	fmt.Println(hasAlternateBits(num))
}

func hasAlternateBits(num int) bool {
	for num != 0 {
		prev := num % 2 // or prev := num & 1
		num = num >> 1
		curr := num % 2 // or curr := num & 1
		num = num >> 1

		if prev == curr {
			return false
		}
	}
	return true
}
