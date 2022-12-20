package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abc"
	permutations(s, 0, len(s))
	fmt.Println("Algorithm2...")
	var ds []string
	visited := make([]bool, len(s))
	permutations2(s, ds, visited)
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

func permutations2(s string, ds []string, visited []bool) {

	if len(ds) == len(s) {
		fmt.Println(strings.Join(ds, ""))
		return
	}

	for i := 0; i < len(s); i++ {
		if !visited[i] {
			visited[i] = true
			ds = append(ds, string(s[i]))
			permutations2(s, ds, visited)
			ds = ds[:len(ds)-1]
			visited[i] = false
		}
	}
}

func swap(s string, i, j int) string {
	sArr := []rune(s)
	sArr[i], sArr[j] = sArr[j], sArr[i]
	return string(sArr)
}
