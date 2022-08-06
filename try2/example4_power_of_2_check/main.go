package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOf2(5))
	fmt.Println(isPowerOf2(10))
	fmt.Println(isPowerOf2(13))
}

func isPowerOf2(num int) bool {
	var r int
	for i := 1; r < num; i++ {
		r = i * 2
	}
	return r == num
}
