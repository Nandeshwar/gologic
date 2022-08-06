package main

import (
	"fmt"
)

func main() {
	fmt.Println(countOf1(7))
	fmt.Println(countOf1(8))

	fmt.Println("What is at 3rd bit")
	fmt.Println(whatIsAt3rdBit(16))
	fmt.Println(whatIsAt3rdBit(7))

	fmt.Println("-------- how many bits required to change from 1 number to another. ex 14 to 15")
	howManyBitsToChange14To15()

}

func countOf1(num int) int {
	cnt := 0
	for num != 0 {
		cnt++
		num = num & (num - 1)
	}
	return cnt
}

func whatIsAt3rdBit(num int) int {
	mask := 1 << 2 // left shit of 1 by n-1
	num = num & mask
	return num // non zero will be 1
}

// or operation with mask will set bit

// How to find how many bit is required to change from a to b
// xor  r := a ^ b
// then count how many 1

func howManyBitsToChange14To15() {
	a := 14 // 1110
	b := 15 // 1111

	r := a ^ b

	cnt := 0
	for r != 0 {
		cnt++
		r = r & (r - 1)
	}

	fmt.Println("count=", cnt) // expection is 1

}
