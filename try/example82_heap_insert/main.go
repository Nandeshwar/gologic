package main

import (
	"fmt"
)

/*
1. heap should be full binary tree
2. there should be at most 2 leaf nodes
3. empty left node should not be allowed

how to know parent from any node:
  i / 2
how to go to left child
i * 2
how to go to right child
(i * 2) + 1

max heap and min heap
max heap: every child node should be less than parent


           20
        18     19
     15   17

    start index:
      1 -- 20
      2 -- 18
      3 -- 19
      4 -- 15
      5 -- 17

    parent of 19: 3 / 2 = 1
    parent of 17: 5 / 2 = 2

    left child of 18 = 2 * 2 = 4

    insert logic:
       Always insert item in the last. Lets insert 25

            20
        18         19
     15   17    25

    next phase
     		20
        18         25
     15   17    19

    next phase

    			25
        18         20
     15   17    19
*/
func heap_insert(a *[]int, item int) {

	*a = append(*a, item)

	parentIndex := len(*a) / 2
	currentIndex := len(*a) - 1
	for parentIndex > 0 {
		if (*a)[currentIndex] < (*a)[parentIndex] {

			return
		}

		(*a)[currentIndex], (*a)[parentIndex] = (*a)[parentIndex], (*a)[currentIndex]
		currentIndex = parentIndex
		parentIndex = currentIndex / 2
	}

}

// Delete strategy
// Delete from root
// bring last node to root
// compare if root node is less than left or right node. If so then swap

/*
 			20
        18     19
     15   17


    			17
        18     19
     15

     		19
        18     17
     15

*/
func heap_delete(a *[]int) {
	(*a)[1] = (*a)[len(*a)-1]
	(*a)[len(*a)-1] = 0
	currentIndex := 1

	leftIndex := 2 * currentIndex
	rightIndex := (2 * currentIndex) + 1

	for currentIndex < len(*a)-1 {
		fmt.Println("currentIndex: ", currentIndex, "leftIndex=", leftIndex, "rightIndex=", rightIndex)

		if (*a)[leftIndex] > (*a)[rightIndex] && (*a)[currentIndex] < (*a)[leftIndex] {
			(*a)[currentIndex], (*a)[leftIndex] = (*a)[leftIndex], (*a)[currentIndex]
			currentIndex = leftIndex
		} else if (*a)[rightIndex] > (*a)[leftIndex] && (*a)[currentIndex] < (*a)[rightIndex] {
			(*a)[currentIndex], (*a)[rightIndex] = (*a)[rightIndex], (*a)[currentIndex]
			currentIndex = rightIndex
		}

		leftIndex = 2 * currentIndex
		rightIndex = (2 * currentIndex) + 1

		if leftIndex > len(*a)-1 && rightIndex >= len(*a)-1 {
			break
		}
	}

}

func create_heap(a *[]int) {
	// a := []int{20, 18, 19, 15, 17}
	b := []int{0}
	b = append(b, *a...)
	*a = b
}
