package main

import "math"

/*
Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
*/
func main() {

}

func minWindow(s string, t string) string {

	if len(t) > len(s) {
		return ""
	}

	mt := make(map[byte]int)
	ms := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		mt[t[i]] = mt[t[i]] + 1

	}

	resultCnt := math.MaxInt
	resultIndex1 := -1
	resultIndex2 := -1

	l := 0
	for r := 0; r < len(s); r++ {
		if _, ok := mt[s[r]]; ok {
			ms[s[r]] = ms[s[r]] + 1
		}
		// delete letter from 1st map as long as both map has same value
		for MapEqual(ms, mt) {

			tmpResultCnt := r - l + 1

			if tmpResultCnt < resultCnt {
				resultCnt = tmpResultCnt
				resultIndex1 = l
				resultIndex2 = r

			}

			if l >= len(s) {
				break
			}

			if cnt, ok := ms[s[l]]; ok {
				if cnt <= 1 {
					delete(ms, s[l])
				} else {
					ms[s[l]] = ms[s[l]] - 1
				}

			}
			l++
		}

	}
	if resultIndex1 == -1 {
		return ""
	}
	return s[resultIndex1 : resultIndex2+1]
}

func MapEqual(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		v2, ok := m2[k]
		if !ok {
			return false
		}

		if v < v2 {
			return false
		}
	}
	return true
}
