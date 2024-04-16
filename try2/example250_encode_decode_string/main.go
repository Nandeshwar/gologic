package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello how are you doing"
	// output: should be same after decoding

	encodedStr := encode(str)
	fmt.Println("encoded str=", encodedStr)
	decodedStr := decode(encodedStr)
	fmt.Println("decoded str=", decodedStr)
}

func encode(str string) string {
	var encodedStr string
	for _, word := range strings.Split(str, " ") {
		encodedStr += strconv.Itoa(len(word)) + "#" + word
	}
	return encodedStr
}

func decode(str string) string {
	var decodedStr string
	for i := 0; i < len(str); {
		// get word len value till #: Ex: 5#Hello : want to extract 5
		j := i
		for str[j] != '#' {
			j++
		}
		wordLen := str[i:j]
		wordLenInt, _ := strconv.Atoi(wordLen)

		word := str[j+1 : j+1+wordLenInt]
		decodedStr += word + " "
		i = j + 1 + wordLenInt
	}
	return decodedStr[0:strings.LastIndex(decodedStr, " ")]
}

/*
output:
 encoded str= 5#hello3#how3#are3#you5#doing
decoded str= hello how are you doing

*/
