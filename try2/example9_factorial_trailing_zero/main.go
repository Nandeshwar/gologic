package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorialTrailingZero(20))
}

func factorialTrailingZero(n int) int {
	result := 0
	for i := 5; i <= n; i *= 5 {
		result += n / i
	}
	return result
}
