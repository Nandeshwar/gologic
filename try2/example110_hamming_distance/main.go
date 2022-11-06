package main

import (
	"fmt"
)

/*
hamming distance is:
number bits that are different in position
ex:
  1 - 0001
  4 - 0100
 number of bits that are different count are: 2

solutin:
   a ^ b  --- same bit will be 0 and different bit will be 1

   0001
 ^ 0100
--------
   0101
      then count number of 1 in the result


*/
func main() {
	a := 1
	b := 4

	d1 := hammingDistance(a, b)
	fmt.Println("d1=", d1)

	a = 9
	b = 14
	d2 := hammingDistance(a, b)
	fmt.Println("d2=", d2)

	a = 4
	b = 8
	d3 := hammingDistance(a, b)
	fmt.Println("d3=", d3)
	fmt.Println("======================")

	a = 1
	b = 4

	d1 = hammingDistance2(a, b)
	fmt.Println("d1=", d1)

	a = 9
	b = 14
	d2 = hammingDistance2(a, b)
	fmt.Println("d2=", d2)

	a = 4
	b = 8
	d3 = hammingDistance2(a, b)
	fmt.Println("d3=", d3)
}

func hammingDistance(a, b int) int {
	count := 0

	n := a ^ b

	for n != 0 {
		count++
		n = n & (n - 1)
	}
	return count
}

func hammingDistance2(a, b int) int {
	c := a ^ b
	cnt := 0
	for i := 0; i < 32; i++ {
		mask := 1 << i
		if c&mask != 0 {
			cnt++
		}
	}
	return cnt
}
