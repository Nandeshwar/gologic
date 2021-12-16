package main

//1. check if 1st item < mid that make sures left items are sorted
// 2. now check if search item < mid  
 //  3. if fo, check if search item is between beg to end , if so end = mid -1 else beg = mid + 1

// if 1st item not < mid 
// then check if search item > mid and less than end then make beg = end + 1 else end mid -1
func searchInSortedCircularArr(arr []int, item int) int {
	beg := 0
	end := len(arr) - 1

	for beg <= end {
		mid := beg + (end-beg)/2
		if arr[mid] == item {
			return mid
		} else if arr[beg] < arr[mid] {
			if item > arr[beg] && item < arr[mid] {
				end = mid - 1
			} else {
				beg = mid + 1

			}
		} else {
			if item > arr[mid] && item < arr[end] {
				beg = end + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
