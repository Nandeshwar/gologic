package main

import (
	"fmt"
)

func permutation(s string, left, right int, allStrList []string) []string {
	if left == right {
		allStrList = append(allStrList, s)
		fmt.Println(s)
		fmt.Println(allStrList)
		return allStrList
	}
	for i := left; i <= right; i++ {
		strArr := []rune(s)
		strArr[left], strArr[i] = strArr[i], strArr[left]
		allStrList = permutation(string(strArr), left+1, right, allStrList)
		strArr[left], strArr[i] = strArr[i], strArr[left]
	}
	return allStrList
}
