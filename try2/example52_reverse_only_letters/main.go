package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(reverseOnlyLetters("ra##d#abc"))
}

func reverseOnlyLetters(str2 string) string {

	str := []rune(str2)
	i := 0
	j := len(str) - 1

	for i < j {
		if !unicode.IsLetter(str[i]) {
			i++
			continue
		}
		if !unicode.IsLetter(str[j]) {
			j--
			continue
		}

		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	return string(str)
}
