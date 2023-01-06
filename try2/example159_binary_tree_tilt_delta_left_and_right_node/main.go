package main

import (
	"fmt"
)

/*
        5
      3    6
    2  4

   output:
total= 2 abs(left-right)= 0 tilt= 0
total= 4 abs(left-right)= 0 tilt= 0
total= 9 abs(left-right)= 2 tilt= 2
total= 6 abs(left-right)= 0 tilt= 2
total= 20 abs(left-right)= 3 tilt= 5
tilt= 5


how
  total
        20
      9    6    : left node difference 2 and total 9
    2  4

tilt: diff
       3      difference is there
      2   0   difference is here
    0  0
*/

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func main() {
	two := &Tree{item: 2}
	six := &Tree{item: 6}
	four := &Tree{item: 4}
	three := &Tree{item: 3, left: two, right: four}
	root := &Tree{item: 5, left: three, right: six}

	var tilt int
	findTilt(root, &tilt)
	fmt.Println("tilt=", tilt)
}

func findTilt(root *Tree, tilt *int) int {
	if root == nil {
		return 0
	}

	left := findTilt(root.left, tilt)
	right := findTilt(root.right, tilt)

	total := root.item + left + right
	*tilt += abs(left - right)

	fmt.Println("total=", total, "abs(left-right)=", abs(left-right), "tilt=", *tilt)
	return total
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
