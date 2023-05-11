package main

/*
You are given a string s. We want to partition the string into as many parts as possible so that each
letter appears in at most one part.

Note that the partition is done so that after concatenating all the parts in order,
the resultant string should be s.

Return a list of integers representing the size of these parts.



Example 1:

Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.
Example 2:

Input: s = "eccbbbbdec"
Output: [10]

Algorithm:
1. store letter and last index
   size = 0
   end = 0

2. iterate loop
   size = ++
   end = max(m[c], end)

   if i == end {
	 result += size
	size = 0

*/

import (
	"fmt"
)
func main() {
	s := "ababcbacadefegdehijhklij"
	m := make(map[rune]int)
	// add character and it's last index
	for i, c := range s {
		m[c] = i
	}

	size := 0
	end := 0

	var result []int
	for i, c := range s {
		size++
		
		// update end if last index of any character is more 
		// and also keep cheking it's last charater comparing with i then that's a partition neetcode.
		end = max(m[c], end)
		
		// once i and end same, that means last character and that's a partition - n
		if i == end {
			result = append(result, size)
			size = 0
		}
		

	}
	fmt.Println("result=", result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
