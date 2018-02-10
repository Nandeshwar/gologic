package main

import "fmt"

func main() {
	fmt.Println(fact(5, 1))
}

func fact(num int, result int) int{
	if num  <= 0 {
		return result
	}
	result = num * result
	num--
	return fact(num, result)
}
