package main

import (
	"testing"
)

func BenchmarkFactorial(b *testing.B) {

	for n := 0; n < b.N; n++ {
		factorial(5)
	}
}

func BenchmarkFactorialWithoutRecusion(b *testing.B) {

	for n := 0; n < b.N; n++ {
		factorialWithoutRecursion(5)
	}

	b.ReportAllocs()
}
