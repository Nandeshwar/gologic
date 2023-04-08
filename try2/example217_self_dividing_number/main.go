package main

import (
	"fmt"
)

func main() {
	/*
			A self-dividing number is a number that is divisible by every digit it contains.

		For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.
		A self-dividing number is not allowed to contain the digit zero.

		Given two integers left and right, return a list of all the self-dividing numbers in the range [left, right].



		Example 1:

		Input: left = 1, right = 22
		Output: [1,2,3,4,5,6,7,8,9,11,12,15,22]
	*/

	fmt.Println(selfDividingNumbers(1, 22))

}

func selfDividingNumbers(left, right int) []int {
	result := []int{}
	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			result = append(result, i)
		}
	}
	return result
}

func isSelfDividingNumber(num int) bool {
	if num == 0 {
		return false
	}

	originalNum := num
	for num > 0 {
		remainder := num % 10
		if remainder == 0 {
			return false
		}

		if originalNum%remainder != 0 {
			return false
		}

		num = num / 10
	}
	return true
}
