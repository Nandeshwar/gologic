/*
https://www.youtube.com/watch?v=gYmWHvRHu-s

inputBook = 10 20 30 40
inputStudent = 2
assign book to each students in all possible way and find max
Then find mix of max
eg. s1 - 10
    s2 - 20 + 30 + 40 = 80
         max 80

    s1 - 10 + 20 = 30
    s2 - 30 + 40 = 70
        max 70

    s1 = 10 + 20 + 30 = 60
    s2 = 40
       max 60

    min of max - 80, 70, 60   is 60
     so Ans is 60

 How to solve?
 use binary search
low := a[0]
end := sum of all items in arr
result = -1


find mid
   if canAllocate(mid) for 2 student
      save mid as result and again for minimum value
      search  between beg to end = mid -1 and check mid if canAllocate numbers.....allocate logic...run through loop
                                                                                and check if items can be be allocated in the mid
                                                                                  if not return false
   else
     search from beg= mid+1 to end


*/
package main

import (
	"github.com/logic-building/functional-go/fp"
)

func minOfMaxBookAllocation(a []int, student int) int {
	beg := a[0]
	end := fp.ReduceInt(func(a, b int) int {
		return a + b
	}, a)

	result := 0

	for beg <= end {
		mid := beg + (end-beg)/2
		//mid := (beg + end) / 2
		if canAllocate(mid, a, student) {
			result = mid
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}
	return result
}

func canAllocate(mid int, a []int, student int) bool {
	page := 0
	allocatedStudent := 1

	for _, item := range a {
		if item > mid {
			return false
		}

		if (item + page) > mid {
			allocatedStudent++
			page = item
			if allocatedStudent > student {
				return false
			}
		} else {
			page += item
		}

	}

	return true
}
