package main

import "fmt"

type AnyFun func (a int, b int) int
func main(){
	myFun(add)
	myFun(sub)
}

func myFun(f AnyFun){
	fmt.Println(f(20, 10))
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
