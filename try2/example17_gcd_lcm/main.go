package main

import (
	"fmt"
)

func main() {
	fmt.Println(findGcd1(8, 12))  // expectation 4
	fmt.Println(findGcd11(8, 12)) // expectation 4
	fmt.Println(findGcd1(5, 7))   // expectation 1

	fmt.Println(findLcm(8, 12)) // expectation 24

	// experimental other approach
	fmt.Println("experimental others")

	fmt.Println(findGcd2(8, 12))
	fmt.Println(findGcd2(12, 8))
	fmt.Println(findGcd3(12, 8)) // brute force
	fmt.Println(findLcm2(8, 12)) // expectation 24 -- brute force

}

// Euclid's method
func findGcd1(a, b int) int {
	if b == 0 {
		return a
	}

	return findGcd1(b, a%b)
}

func findLcm(a, b int) int {
	return (a * b) / findGcd1(a, b)
}

func findGcd2(a, b int) int {
	if b == 0 {
		return a
	}
	if b > a {
		a = a ^ b
		b = a ^ b
		a = a ^ b
	}
	return findGcd2(a-b, b)
}

func findGcd11(a, b int) int {
	for a != 0 && b != 0 {
		if b != 0 {
			a = a % b
		}
		if a != 0 {
			b = b % a
		}
	}
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return 0
}

// brute force
// divide by 2, 3, 4 , 5 as long as it is less than both number. and highest number that divides will be GCD
func findGcd3(a, b int) int {
	gcd := 1
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return gcd
}

// brute force
// 12 and 8
// consider 12 is lcm
// divide lcm by 12 and 8 if divisible then that's lcm if not increment lcm by 1
func findLcm2(a, b int) int {
	lcm := a
	if b > a {
		lcm = b
	}

	for {
		if lcm%a == 0 && lcm%b == 0 {
			break
		}
		lcm++
	}
	return lcm
}
