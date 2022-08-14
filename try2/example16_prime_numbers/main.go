package main

import "fmt"

const N = 100

var primeArr [N + 1]bool

func main() {
	fmt.Println(isPrime1(17))
	primePreprocessing()
	fmt.Println(isPrime2(11))
}

func isPrime1(num int) bool {

	if num == 1 || num == 2 {
		return true
	}
	if num <= 0 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num&1 == 0 {
			return false
		}
	}
	return true
}

func isPrime2(num int) bool {
	if num <= 0 {
		return false
	}
	return primeArr[num]
}

func primePreprocessing() {
	for i := 1; i <= 100; i++ {
		primeArr[i] = true
	}

	for i := 2; i*i <= N; i++ {
		for j := i << 2; j <= N; j += i {
			primeArr[j] = false
		}
	}
}
