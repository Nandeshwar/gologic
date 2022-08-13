package main

import (
	"fmt"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	var result []string

	for i := 1; i <= 15; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, fmt.Sprintf("%d", i))
		}
	}
	return result
}
