package main

import (
	"fmt"
	"unicode"
)

func main() {
	//str := "ram sita t is Ma r"
	str := "A man, a plan, a canal: Panama"
	//str := "m ada m"
	//str := "ram"
	fmt.Println(isPalindrome(str))
}

func isPalindrome(str string) bool {
	i := 0
	j := len(str) - 1

	for i < j {
		if !unicode.IsLetter(rune(str[i])) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(str[j])) {
			j--
			continue
		}

		if i < j && unicode.ToLower(rune(str[i])) != unicode.ToLower(rune(str[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
