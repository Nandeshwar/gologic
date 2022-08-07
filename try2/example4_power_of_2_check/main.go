package main

import (
	"fmt"
)

func main() {
	fmt.Println("is multiple of 2......")
	fmt.Println(isMulitplieOf2(5))
	fmt.Println(isMulitplieOf2(10))
	fmt.Println(isMulitplieOf2(13))

	fmt.Println("is Power of 2")
	fmt.Println(isPowerOf2(5))
	fmt.Println(isPowerOf2(10))
	fmt.Println(isPowerOf2(13))
	fmt.Println(isPowerOf2(16))
}

func isMulitplieOf2(num int) bool {
	var r int
	for i := 1; r < num; i++ {
		r = i * 2
	}
	return r == num
}

func isPowerOf2(n int) bool {
	i := 1
	for i < n {
		i *= 2
	}
	return i == n
}
