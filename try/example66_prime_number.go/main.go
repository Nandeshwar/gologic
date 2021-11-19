package main

func primeNumberBetween(begin, end int) []int {
	var primeNumbers []int
	for num := begin; num <= end; num++ {
		if checkPrime(num) {
			primeNumbers = append(primeNumbers, num)
		}
	}
	return primeNumbers
}

func checkPrime(num int) bool {
	if num <= 0 || num == 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primeNumberBetweenAlgorith2(begin, end int) []int {
	var primeNumbers []int

	// boolean array size will be 1 more than length
	// Ex: 5
	// 0 1 2 3 4 5  ---- Total 6 so that if I say 2 it will go 3rd index i.e 2
	primeNumbersBool := make([]bool, end+1)
	fillPrimeNumberBoolWithTrue(primeNumbersBool)

	primeNumbersBool[0] = false
	primeNumbersBool[1] = false

	// Fill primeNumbersBool to false for multile of each number
	// say 2: multiple of 2 is 4, 6   will be filled false because it is not prime number
	// say 3: multple of 3 is 6 9 12 will be filled with false as this is not prime number
	for i := 2; i*i <= end; i++ {
		for j := i * 2; j <= end; j += i {
			primeNumbersBool[j] = false
		}
	}

	for num := begin; num <= end; num++ {
		if primeNumbersBool[num] {
			primeNumbers = append(primeNumbers, num)
		}
	}
	return primeNumbers
}

func fillPrimeNumberBoolWithTrue(primeNumbersBool []bool) {
	for i := range primeNumbersBool {
		primeNumbersBool[i] = true
	}
}
