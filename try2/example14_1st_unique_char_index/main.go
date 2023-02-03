package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findFirstUniqueLetter("madame"))
	fmt.Println(findFirstUniqueLetter2("madame"))
}

func findFirstUniqueLetter(str string) int {
	m := map[rune]int{}
	for i, c := range str {
		_, ok := m[c]
		if ok {
			m[c] = -1
		} else {
			m[c] = i
		}
	}

	min := int(^uint(0) >> 1)
	fmt.Println("max int value=", min)

	fmt.Println(^uint(0))
	for _, v := range m {
		if v == -1 {
			continue
		}
		min = int(math.Min(float64(v), float64(min)))
	}
	return min
}

func findFirstUniqueLetter2(str string) int {
	a := [26]int{}
	for i := 0; i < 26; i++ {
		a[i] = -1
	}

	for ind, v := range str {

		if a[v-'a'] > -1 {
			a[v-'a'] = -1
		} else {
			a[v-'a'] = ind
		}

	}
	fmt.Println(a)

	min := int(^uint(0) >> 1)
	for i := 0; i < 26; i++ {
		if a[i] == -1 {
			continue
		}
		min = Min(min, a[i])
	}
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
