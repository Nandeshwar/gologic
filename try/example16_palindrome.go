package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("Madam"))
	fmt.Println(isPalindrome("Madam1"))
}

func isPalindrome(s string) bool {
	str := []rune(strings.ToLower(s))

	var str2 []rune

	for _, v := range str {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			str2 = append(str2, v)
		}
	}

	i := 0
	j := len(str2) - 1

	for i < j {
		if str2[i] != str2[j] {
			return false
		}
		i++
		j--
	}
	return true
}