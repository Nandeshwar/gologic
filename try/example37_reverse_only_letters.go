package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "abc#de"

	strChr := []rune(str)

	i := 0
	j := len(strChr) - 1

	for i < j {
		if !unicode.IsLetter(strChr[i]) {
			i++
			continue
		}

		if !unicode.IsLetter(strChr[j]) {
			j--
			continue
		}
		strChr[i], strChr[j] = strChr[j], strChr[i]
		i++
		j--
		
	}

	fmt.Println(string(strChr))
}