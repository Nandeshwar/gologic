package main

import (
	"fmt"
)

// diameter = max(leftHeight + rightHeight)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func main() {
	/*
					10
				5       9
			3      6

			     diameter = 3

				10
		     5
		  3	   6
		2        7
		           8

				diameter = 5
	*/
	nine := &Tree{item: 9}
	six := &Tree{item: 6}
	three := &Tree{item: 3}
	five := &Tree{item: 5, left: three, right: six}
	root1 := &Tree{item: 10, left: five, right: nine}

	eight2 := &Tree{item: 8}
	seven2 := &Tree{item: 7, right: eight2}
	two2 := &Tree{item: 2}
	six2 := &Tree{item: 6, right: seven2}
	three2 := &Tree{item: 3, left: two2}
	five2 := &Tree{item: 5, left: three2, right: six2}
	root2 := &Tree{item: 10, left: five2}

	fmt.Println("Diameter of binary tree")
	var diameter1 int
	var diameter2 int
	findDiameter(root1, &diameter1)
	findDiameter(root2, &diameter2)
	fmt.Println("diameter1=", diameter1)
	fmt.Println("diameter2=", diameter2)

}

// function returns height and stroe the diameter in reference variable
func findDiameter(root *Tree, diameter *int) int {
	if root == nil {
		return 0
	}

	leftHeight := findDiameter(root.left, diameter)
	rightHeight := findDiameter(root.right, diameter)

	*diameter = max(*diameter, leftHeight+rightHeight)

	return max(leftHeight, rightHeight) + 1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
