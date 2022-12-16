package main

import (
	"fmt"
)

func main() {
	// check if string is palindrome, you may delete 1 character
	// str := "eye" // : true
	//str := "eyee" // : true
	//str := "madam1" // true
	str := "madam12" // false

	fmt.Println(isPalindrome(0, len(str)-1, str))
}

func isPalindrome(i, j int, str string) bool {
	for i < j {
		if str[i] != usstr[j] {
			return isPalindrome2(i+1, j, str) || isPalindrome2(i, j-1, str)
		}

		i++
		j--
	}
	return true
}

func isPalindrome2(i, j int, str string) bool {
	for i < j {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}
	return true
}
