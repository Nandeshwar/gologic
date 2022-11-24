package main

import (
	"fmt"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func main() {

	/*
					1
				0		 1
			  0   1         1

			output: 16
			100 + 101 + 111
			 4 + 5 + 7 = 16


			step1: first node
			    sum = 0
				    left shift by 1
				    0
					+ 1
					result= 1
					will be passed

		 	step2: 2nd node left side
			   sum = 1
			      left << by
				  10
				   + 0 (2nd node)
				result := 10 ~~ 2

			step3: 3rd node left side
			    sum = 2
				left shift by 1
				 100
				   + 0(3rd node)
				result = 100 ~~ 4

				now left and right null
				return 4
	*/
	zero := &Tree{item: 0, left: &Tree{item: 0}, right: &Tree{item: 1}}
	one := &Tree{item: 1, right: &Tree{item: 1}}

	root := &Tree{item: 1, left: zero, right: one}
	fmt.Println(sumOfLeaveNodes(root, 0))
}

func sumOfLeaveNodes(t *Tree, sum int) int {
	if t == nil {
		return 0
	}

	sum = sum << 1
	sum += t.item

	if t.left == nil && t.right == nil {
		return sum
	}

	// Above logic. do for 1st node and retrun sum and believe pass sum and process for left and right

	leftSum := sumOfLeaveNodes(t.left, sum)
	rightSum := sumOfLeaveNodes(t.right, sum)

	return leftSum + rightSum
}
