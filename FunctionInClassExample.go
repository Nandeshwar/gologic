package main

import "fmt"

type MyClass struct {
	id int
	myFun newFun

}

type newFun func (int, int) (int)

func main(){
	myClass := MyClass {
		id : 1,
		myFun:  sum,
	}

	fmt.Println(myClass.id)
	fmt.Println(myClass.myFun(10, 20))
	myClass.add()
}

func sum(a int, b int) int {
	return a + b
}

func(myClass MyClass) add(){
	fmt.Println(myClass.id)
}