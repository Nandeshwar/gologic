package main

/*
Some notes:
  How to find bits  for certain number such as 13 - 1101
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
     011


     100  ---- reverse
      +1 ---- add 1
     ---
     101


     111
     101
    ---
     1 0 0  -- 4

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


mask:
  1 << nth position

&: get bit
   number & mask --- result non zero will say bit 1
                     if result zero then zero

 | : set bit
     number | mask --- set bit to 1

*/

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(swap(10, 20))
	fmt.Println(isEven(10))
	fmt.Println(isEven(11))
	fmt.Println(multiplyBy2(10))

	fmt.Println("Single non repeated number where other numbers repeat by 2")
	fmt.Println(singleNonRepeatedNumber([]int{1, 2, 1, 2, 3, 4, 5, 4, 5})) // expected answer 3
	fmt.Println("Two non repeated number where other numbers repeat by 2")
	fmt.Println(twoNonRepeatedNumber([]int{1, 2, 1, 2, 3, 4, 6, 4})) // expectation 3 and 5

	fmt.Println("Single non repeated number where all other numbers repated by 3")
	fmt.Println(singleNonRepeatedWhereOtherRepeatedBy3([]int{1, 2, 3, 4, 1, 2, 3, 1, 2, 3})) // expection is 4
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

// xor properties

/*
1. same number ^ same number = 0
2. any number ^ 0 = same number
*/

// program to find single non repeated nubmer where other numbers repeat 2 times

func singleNonRepeatedNumber(nArr []int) int {
	result := 0
	for _, v := range nArr {
		result ^= v
	}
	return result
}

// get two non repeated numbers where other numbers repeat by 2
func twoNonRepeatedNumber(nArr []int) (int, int) {
	result := 0
	var evenArr []int
	var oddArr []int

	for _, v := range nArr {
		if v&1 != 0 {
			oddArr = append(oddArr, v)
		} else {
			evenArr = append(evenArr, v)
		}

		result ^= v
	}

	num1 := result
	for _, v := range evenArr {
		num1 ^= v
	}

	num2 := result
	for _, v := range oddArr {
		num2 ^= v
	}

	return num1, num2
}

/*
	1. create 32 bit count arry
 	2. find each bit is 1 or not for all the items in array. if 1 increment the counter
	3. if counter value is multiple of 3, then make it 0 otherwise 1
	4. convert counter array(binary value) to decimal and that will be the result
*/
func singleNonRepeatedWhereOtherRepeatedBy3(nArr []int) int {
	countArr := make([]int, 32)
	for i := 0; i < 32; i++ {
		mask := 1 << i

		for _, v := range nArr {
			r := v & mask
			if r != 0 {
				countArr[31-i]++
			}
		}

		// turn counter to 0 if divisible by 3 otherwise 1
		if countArr[i]%3 == 0 {
			countArr[i] = 0
		} else {
			countArr[i] = 1
		}
	}

	fmt.Println(countArr)

	// convert bits string to decimal
	var bitsStr string
	for _, v := range countArr {
		bitsStr += fmt.Sprintf("%d", v)
	}
	decimalNo, _ := strconv.ParseInt(bitsStr, 2, 64)

	return int(decimalNo)
}
