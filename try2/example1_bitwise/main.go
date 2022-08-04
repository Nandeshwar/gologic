package main

/*
Some notes:
  How to find bit for certain number
  1. divide number by 2 and put the remainder on side. Take the numbers from botton to top
  2. 13
    it is odd so    1
    divide by 2
     6 it's even   0 1
     3 its odd   1 0 1
     1

    so 1101 is 13


    Binary Add
     1 + 1 = 0 then carry over

    1 1 0 1
    0 1 0 1
    ---------
   1 0  0  1   0

   Binary Subtraction
   7 - 3
   find 2's complement of 3
     11


     00  ---- reverse
     +1 ---- add 1
     ---
     10


     111
     +10
    ---
    1 0 1  -- 5

    Binary xor
     1 1 ----0
     0 0 ---- 0
     1 0 -- 1
     0 1 --1

  Swap number:
    a = a ^ b
    b = a ^ b
    a = a ^ b

  multiply number by 2
  result = 10 << 1

  Divide number by 2
  result = 10 >> 1

  find even or odd
-------------
  a & 1  == 0 is even

*/

import (
	"fmt"
)

func main() {
	fmt.Println(swap(10, 20))
	fmt.Println(isEven(10))
	fmt.Println(isEven(11))
	fmt.Println(multiplyBy2(10))
}
func swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func isEven(n int) bool {
	return n&1 == 0
}

func multiplyBy2(n int) int {
	return n << 1
}
