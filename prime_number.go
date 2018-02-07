package main

import "fmt"

func main(){
	fmt.Println(isPrime(3))
}

func isPrime(num int) bool {
	if num == 1 || num == 2 {
		return true
	}

	for i := 2; i < num / 2 + 1; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}
