package main

import (
	"fmt"
)

func main() {
	fmt.Println(powerOfN(2, 4))
	fmt.Println(power(2, 4))
}

// o(n)
func powerOfN(x, n int) int {
	if n == 0 {
		return 1
	}
	return x * powerOfN(x, n-1)
}

// o(log(n))
func power(x, n int) int {
	if n == 0 {
		return 1
	}
	r := power(x, n/2) * power(x, n/2)

	if n%2 != 0 {
		r = r * x
	}
	return r
}
