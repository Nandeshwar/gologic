package main

import (
	"fmt"
)

func main() {
	s := "abc"
	permutations(s, 0, len(s))
}

func permutations(s string, left, right int) {

	if left == right {
		fmt.Println(s)
		return
	}

	for i := left; i < right; i++ {
		s = swap(s, i, left)
		permutations(s, left+1, right)
		s = swap(s, i, left)
	}
}

func swap(s string, i, j int) string {
	sArr := []rune(s)
	sArr[i], sArr[j] = sArr[j], sArr[i]
	return string(sArr)
}
