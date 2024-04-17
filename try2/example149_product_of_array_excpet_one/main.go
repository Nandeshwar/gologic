package main

import (
	"fmt"
)

/*
strategy to solve:

array  =   1   2    3   4
prefix =   1   2    6   24 (multiply: start from beginning of array)
postfix=   24  24  12   4 (multiply: start from end of array

1st pass prepare prefix
1st pass:

	 prefix * post fix
	 for 1 no prefix i.e. 1 but post fix is 24 = so 24
	 for 2 prefix(1) * postfix(12)
	 for 3 prefix(2 * 4)
	 for 4 prefix( 6 and no postfix(i.e 1)
	24 12 8 6

strategy2:
--------
No extra memory: use 1 output array

store prefix value just below of the array in output
array:    1    2   3   4
output:   1    1   2   6

	             6 : because 4 prefix is 6 and no postfix
	          8    : 3 prefix is 2 * post prefix(4)
	     12         : 2 prefix 1 and post prefix( 3 * 4 = 12)
	24              : 1 prefix 1 and post prefix( 3 * 4 * 2 = 24

that's why
output:  24   12  8   6

traverse output from end : same logic prefix and postfix: postfix need to be calcuated
*/
func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(a))
	fmt.Println(productExceptSelf2(a))
}

func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))

	l := 1
	out[0] = 1
	for i := 1; i < len(nums); i++ {
		l *= nums[i-1]
		out[i] = l
	}

	r := 1
	for i := len(nums) - 2; i >= 0; i-- {
		r = r * nums[i+1]
		out[i] = out[i] * r
	}

	return out

}

func productExceptSelf2(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	out := make([]int, len(nums))

	left[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i]
	}

	right[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			out[i] = right[i+1]
		} else if i == len(nums)-1 {
			out[i] = left[len(nums)-2]
		} else {
			out[i] = left[i-1] * right[i+1]
		}
	}
	return out
}

/*
output:
left= [1 1 2 6]
right= [24 12 8 6]
[24 12 8 6]

*/
