package main

import (
	"fmt"
)

func main() {
	fmt.Println("is Power of 2")
	fmt.Println(isPowerOf2(5))
	fmt.Println(isPowerOf2(10))
	fmt.Println(isPowerOf2(13))
	fmt.Println(isPowerOf2(16))
}

func isPowerOf2(n int) bool {
	i := 1
	for i < n {
		i *= 2
	}
	return i == n
}
