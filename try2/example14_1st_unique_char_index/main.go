package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findFirstUniqueLetter("madame"))
}

func findFirstUniqueLetter(str string) int {
	m := map[string]int{}
	for i, c := range str {
		_, ok := m[string(c)]
		if ok {
			m[string(c)] = -1
		} else {
			m[string(c)] = i
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
