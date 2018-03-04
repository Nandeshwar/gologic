package main

import "fmt"
var  numList []int
const LEN = 3
var top = -1

func main() {

	push(10)
	push(11)
	push(12)
	push(13)

	fmt.Println(numList)

	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
}


func push(item int){
	fmt.Println("length: " , len(numList))
	if len(numList) == LEN {
		fmt.Println("Stack is full")
		return
	} else {
		top++
		numList = append(numList, item)
		fmt.Println("Item added to stack successfully")
	}
}

func pop() (item int){
	if top == -1 {
		fmt.Println("Stack empty")
		return
	}

	item = numList[top]
	top--
	return
}

