package main

import (
	"fmt"
)

func main() {
	fmt.Println(replaceDotWithSquareBrackets("a.b.c")) // expected output: a[.]b[.]c
}

func replaceDotWithSquareBrackets(str string) string {
	newStr := ""
	for _, v := range str {
		strV := string(v)
		if strV == "." {
			newStr += "[.]"
		} else {
			newStr += string(v)
		}
	}
	return newStr
}
