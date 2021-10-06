package main

import (
	"fmt"
)

func main() {
	s1 := "abcdef"

	for i, _ := range s1 {
		fmt.Println(s1[i] - 'a')
	}
}