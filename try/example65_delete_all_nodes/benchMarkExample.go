package main

func factorial(n int) int {
	a := 10

	if n == 0 && a == 10 {
		return 1
	}

	return n * factorial(n-1)
}

func factorialWithoutRecursion(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
