package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseStr("Radha"))
}

func reverseStr(str string) string {
	strArr := []rune(str)

	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	return string(strArr)
}
