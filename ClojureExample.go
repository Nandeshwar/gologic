package main

import "fmt"

func adder() func(int) int {
	sum := 0 // this variable will be treated as global variable within closure call
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	f := adder()
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(3))
}

/*
output:
2
5
8
 */