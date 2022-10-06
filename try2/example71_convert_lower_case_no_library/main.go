package main

import (
	"fmt"
)

func main() {
	fmt.Println(toLower("RAM"))
}

func toLower(s string) string {
	sArr := []rune(s)
	for i := 0; i < len(sArr); i++ {
		sArr[i] = sArr[i] + 32
	}
	return string(sArr)
}
