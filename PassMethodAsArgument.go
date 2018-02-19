package main

import "fmt"

type AnyFun func (a int, b int) int

func main(){
	// Passing named method definition in function
	myFun(add)
	myFun(sub)
	myFun2(add)

	// Passing anonymous method definition in function
	myFun2(func(a int, b int) int{
		return a * b
	})
}

func myFun(f AnyFun){
	fmt.Println(f(20, 10))
}

func myFun2(f func(a int, b int) int){
	fmt.Println(f(10, 20))
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
