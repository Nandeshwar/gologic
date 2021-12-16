package main

// inner loop compares two adjacent items, swap based on condition then continue comparing adjacent itmes.
// Largest item will reach to end
// inner loop will do comparision till n - i in each iteration
func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j+1] < nums[j] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}
