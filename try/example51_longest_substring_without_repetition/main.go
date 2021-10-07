package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("What is RAM?")
	str := "pwwwkew"
	fmt.Println(findLongestLenOfSubstring_withoutrRepetition(str))
}

func findLongestLenOfSubstring_withoutrRepetition(str string) int {
	strArr := []rune(str)
	
	j := 0
	
	max := 0

	m := map[rune]struct{}{}
	for ; j < len(strArr); j++ {
		_, ok := m[strArr[j]]
		if !ok {
			m[strArr[j]] = struct{}{}
			max = int(math.Max(float64(len(m)), float64(max)))
		} else {
			m = map[rune]struct{}{}
			m[strArr[j]] = struct{}{}
		}
	}
	return max
}

