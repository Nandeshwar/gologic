package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("Madam"))
	fmt.Println(isPalindrome("Madam1"))
	
	fmt.Println("Palindrome solution using recursion")
	s := "madam"
	fmt.Println(isPalindromerecursive(s, 0, len(s)-1))
	
	s = "madam1"
	fmt.Println(isPalindromerecursive(s, 0, len(s)-1))
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

func isPalindromerecursive(s string, l, r int) bool {
	if l > r {
		return true
	}
	if s[l] != s[r] {
		return false
	}
	return isPalindromerecursive(s, l+1, r-1)
}