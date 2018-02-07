package main

import "fmt"

func main(){
	fmt.Println("Fibonacci series: ")
	fmt.Println(fib(5))
}

func fib(num int) (fib_series []int) {
	var a = -1
	var b = 1
	var c = a + b
	if num <= 1 {
		fib_series = append(fib_series, c)
		return
	}

	counter := 1

	for counter <= num {
		fib_series = append(fib_series, c)
		a = b
		b = c
		c = a + b
		counter++
	}

	return
}