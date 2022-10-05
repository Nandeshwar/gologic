package main

import (
	"fmt"
)

func main() {
	fmt.Println(add("123", "456"))
}

func add(s1, s2 string) string {
	var s3 string

	i := len(s1) - 1
	j := len(s2) - 1

	carry := 0
	for i >= 0 || j >= 0 {

		sum := carry
		if i >= 0 {
			sum += int(s1[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(s2[j] - '0')
			j--
		}

		r := sum % 10
		carry = sum / 10
		s3 += fmt.Sprintf("%d", r)
	}

	s3 += string(carry)
	var s4 string
	for i := len(s3) - 1; i >= 0; i-- {
		s4 += string(s3[i])
	}
	return s4
}
