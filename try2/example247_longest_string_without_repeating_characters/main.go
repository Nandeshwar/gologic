package main

import "fmt"

/*
Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func main() {
	fmt.Println("")
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	l := 0
	m := 0

	ds := make(map[byte]int)
	for r := 0; r < len(s); r++ {

		ind, ok := ds[s[r]]
		if ok {
			for i := l; i <= ind; i++ {
				delete(ds, s[i])
			}

			l = ind + 1
		}

		ds[s[r]] = r
		m = max(len(ds), m)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
