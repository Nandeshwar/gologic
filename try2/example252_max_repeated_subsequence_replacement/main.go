package main

import "fmt"

func main() {
	// find max repeating subsequence, other charcter can be replaced upto 2 times
	fmt.Println(characterReplacement("nandeshwar", 2))
}

func characterReplacement(s string, k int) int {
	m := make(map[byte]int)

	result := 0
	l := 0
	for r := 0; r < len(s); r++ {
		m[s[r]] = m[s[r]] + 1

		if r-l+1-maxRepeated(m) > k {
			m[s[l]] = m[s[l]] - 1
			l++
		}

		result = Max(result, r-l+1)
	}

	return result
}

func maxRepeated(m map[byte]int) int {
	maxValue := -100
	for _, v := range m {
		maxValue = Max(maxValue, v)
	}
	return maxValue
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
