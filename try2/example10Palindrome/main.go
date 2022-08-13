package main

import (
	"fmt"
)

func main() {
	fmt.Println(strPalindrome("madam"))
	fmt.Println(reverse(-123))
}

func reverse(n int) int {
	result := 0
	for n != 0 {
		r := n % 10
		result = result*10 + r
		n = n / 10
	}
	return result
}

func strPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	strArr := []rune(s)

	for i < j {
		if strArr[i] != strArr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
