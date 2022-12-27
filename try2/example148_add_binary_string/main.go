package main

import (
	"fmt"
)

func main() {
	str1 := "1011"
	str2 := "1011"

	// expected output
	// 10110

	result := add(str1, str2)
	fmt.Println("result=", result)
}

func add(str1, str2 string) string {
	var result string
	carry := 0

	i := len(str1) - 1
	j := len(str2) - 1

	for i >= 0 || j >= 0 {
		sum := carry

		if i >= 0 {
			sum += int(str1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(str2[j] - '0')
			j--
		}

		r := sum % 2
		carry = sum / 2

		result = fmt.Sprintf("%d", r) + result
	}
	if carry == 1 {
		result = fmt.Sprintf("%d", carry) + result
	}
	return result
}
