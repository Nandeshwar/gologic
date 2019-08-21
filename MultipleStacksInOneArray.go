package main

import "fmt"

var arrLen = 9
var stackCount = 3

var arr []int
var stackArr []int

// Track Max and min index of each stack
//stack0: 2 : formula n/3 - 1
// stack1:5 : formula 2n/3 -1
// stack2->8 : formula 3n/3 -1
var stackTopMax []int

//stack0: 0 : initial value
// stack1:3 : formula n/3
// stack2->6 : formula 2n/3 - 1
var stackTopMin []int

// start Array index for each stack:
// stack0: 0,
// stack1: n/3 here 9/3 i.e 3
// stack2: 2n/3 here 18/3 i.e 6
var stackArrIndex []int

func main() {
	arr = make([]int, arrLen)
	stackArr = make([]int, stackCount)
	stackArrIndex = make([]int, stackCount)
	stackTopMax = make([]int, stackCount)
	stackTopMin = make([]int, stackCount)

	if stackCount > arrLen {
		fmt.Println("Increase array length or decrease stack count")
	}

	initialIndex := 0
	stackMin := 0

	for i := 0; i < stackCount; i++ {
		stackArrIndex[i] = initialIndex
		stackTopMax[i] = (((i + 1) * arrLen) / stackCount) - 1
		stackTopMin[i] = stackMin

		initialIndex = ((i + 1) * arrLen) / stackCount
		stackMin = ((i + 1) * arrLen) / stackCount
	}

	push(0, 10)
	push(0, 20)
	push(0, 30)
	push(0, 40)
	push(1, 100)
	push(2, 1000)
	printStack()
	pop(2)
	printStack()
	pop(2)
	printStack()
	push(2, 1000)
	push(2, 1001)
	printStack()

}

func push(stackNumber, val int) {

	if stackArrIndex[stackNumber] > stackTopMax[stackNumber] {
		fmt.Printf("\n stack %d is full", stackNumber)
	} else {
		arr[stackArrIndex[stackNumber]] = val
		stackArrIndex[stackNumber]++
		fmt.Printf("\n Item=%d pushed in stack number=%d", val, stackNumber)
	}
}

func pop(stackNumber int) {
	stackArrIndex[stackNumber]--
	if stackArrIndex[stackNumber] < stackTopMin[stackNumber] {
		fmt.Println("Stack is empty")
	} else {
		arr[stackArrIndex[stackNumber]] = 0
		fmt.Println("Item poped from stack number", stackNumber)
	}
}

func printStack() {
	fmt.Println(arr)
}

/*
output:
 Item=10 pushed in stack number=0
 Item=20 pushed in stack number=0
 Item=30 pushed in stack number=0
 stack 0 is full
 Item=100 pushed in stack number=1
 Item=1000 pushed in stack number=2[10 20 30 100 0 0 1000 0 0]
Item poped from stack number 2
[10 20 30 100 0 0 0 0 0]
Stack is empty
[10 20 30 100 0 0 0 0 0]

 Item=1000 pushed in stack number=2
 Item=1001 pushed in stack number=2[10 20 30 100 0 1000 1001 0 0]
*/
