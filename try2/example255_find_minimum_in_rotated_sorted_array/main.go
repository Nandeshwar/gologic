package main

/*
Example 1:

Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.
Example 2:

Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
Example 3:

Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.
*/
func main() {

}

func findMin(nums []int) int {
	min := nums[0]

	beg, end := 0, len(nums)-1

	for beg <= end {
		if nums[beg] < nums[end] {
			min = Min(nums[beg], min)
			return min
		}

		// mid := beg + (end-beg) / 2
		mid := (beg + end) / 2
		min = Min(nums[mid], min)
		if nums[beg] <= nums[mid] {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}

	return min

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
