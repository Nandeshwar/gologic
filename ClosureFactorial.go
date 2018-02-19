package main

import "fmt"

func factorial() func(num int) int{
	fact := 1
	return func(num int) int{
		fact = fact * num
		return fact
	}
}

func main(){
	f := factorial()
	for i := 1; i <= 5; i++ {
		fmt.Println(f(i))
	}
}