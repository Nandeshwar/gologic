package main

import "fmt"

var arr []int
var stackLen = 3
var stackCount = 0

// stack top pointer for different stack
var stackIndex []int

// This is used to check max allowed length in different stack
var stackMax []int

// This is used to check min allowed length in different stack
var stackMin []int

func main() {
	Push(10)
	print()
	Push(20)
	print()
	Push(30)
	print()
	Push(40)
	print()
	Push(50)
	print()
	Push(60)
	print()
	Push(70)
	print()
	Pop()
	print()
	Pop()
	print()
	Pop()
	print()
	Pop()
	print()
}

func Push(item int) {
	// For first stack
	if stackCount == 0 {
		stackCount++
		arr = make([]int, stackCount*stackLen)

		// This is used to check max allowed length in different stack
		stackMax = make([]int, stackCount)

		// This is used to check min allowed length in different stack
		stackMin = make([]int, stackCount)

		// stack top pointer for different stack
		stackIndex = make([]int, stackCount)

		// stackCount=1, stackLen=3 = (1 * 3) -1 = 2
		stackMax[0] = (stackCount * stackLen) - 1

		stackMin[0] = (stackCount * stackLen) - stackLen

		stackMin[0] = 0
		arr[0] = item
		stackIndex[0]++
		fmt.Printf("\n1st item=%d pushed", item)
	} else {

		// loop through all the stacks
		for i := 0; i < stackCount; i++ {

			// stack is full. create new stack
			if stackFull(i) {

				// if 1st stack is full, before creating , check if the is already a stack, if so skip
				if i+1 < stackCount {
					continue
				}
				stackCount++

				// create new stack
				arr2 := make([]int, stackCount*stackLen)
				arr = mapToNewArr(arr, arr2)

				// new max, min, index array for new stack
				newStackMax := make([]int, stackCount)
				stackMax = mapToNewArr(stackMax, newStackMax)

				stackMin2 := make([]int, stackCount)
				stackMin = mapToNewArr(stackMin, stackMin2)

				newStackIndex := make([]int, stackCount)
				stackIndex = mapToNewArr(stackIndex, newStackIndex)

				// stack current pointer for newly create stack. for second stack, stackCount-2, stackLen-3, (2-1) * 3 - 3
				stackIndex[stackCount-1] = (stackCount - 1) * stackLen

				// insert item
				arr[stackIndex[stackCount-1]] = item

				// increment current pointer
				stackIndex[stackCount-1]++

				// max stack - 5 for 2nd stack of length 3
				stackMax[stackCount-1] = (stackCount * stackLen) - 1

				stackMin[stackCount-1] = (stackCount * stackLen) - stackLen
				return
			} else {
				// if stack is not full, insert item in respective stack in array and increment stackIndex for top pointer
				arr[stackIndex[i]] = item
				stackIndex[i]++
				return
			}
		}
	}
}

func stackFull(stackNumber int) bool {
	if stackIndex[stackNumber] > stackMax[stackNumber] {
		return true
	}
	return false
}

func mapToNewArr(arr, arr2 []int) []int {
	for i, v := range arr {
		arr2[i] = v
	}
	arr = arr2
	return arr
}

func print() {
	fmt.Println(arr)
}

func Pop() {
	stackNumber := stackCount - 1
	stackIndex[stackNumber]--
	fmt.Printf("\nstack=%d, stackIndex=%d: ", stackNumber, stackIndex[stackNumber])
	arr[stackIndex[stackNumber]] = 0
	if stackIndex[stackNumber] <= stackMin[stackNumber] {
		//TODO: Reduce size of array
		stackCount--
	}
}

/*
Output:
  1st item=10 pushed[10 0 0]
[10 20 0]
[10 20 30]
[10 20 30 40 0 0]
[10 20 30 40 50 0]
[10 20 30 40 50 60]
[10 20 30 40 50 60 70 0 0]

stack=2, stackIndex=6: [10 20 30 40 50 60 0 0 0]

stack=1, stackIndex=5: [10 20 30 40 50 0 0 0 0]

stack=1, stackIndex=4: [10 20 30 40 0 0 0 0 0]

stack=1, stackIndex=3: [10 20 30 0 0 0 0 0 0]
*/
