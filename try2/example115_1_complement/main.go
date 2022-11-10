package main

import (
	"fmt"
)

func main() {
	a := 10 // 1010 // expected result - 0101(5)

	result := 0

	i := 0
	for a != 0 {
		if a&1 == 0 {
			result += 1 << i
		}

		i += 1
		a >>= 1
	}

	fmt.Println("result=", result)
}
