package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isAnagram("eeys", "eeyt"))
}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	arr := make([]int, 26)

	strArr := []rune(str2)

	for i, v := range str1 {
		if unicode.IsLower(v) {
			arr[v-97]++
		} else {
			arr[v-65]++
		}

		if unicode.IsLower(strArr[i]) {
			arr[strArr[i]-97]--
		} else {
			arr[strArr[i]-65]--
		}
	}
	fmt.Println(arr)
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}
