package main

import (
	"fmt"
	"math"
)

type node23 struct {
	item int
	left *node23
	right *node23
}

func main() {
	/*
	input:
	   1
	3     2
	5     3  9
	output:
	   1, 3, 9 
	*/

	node5 := node23 {
		item: 5,
	}

	node3 := node23 {
		item: 3,
	}

	node9 := node23 {
		item: 9,
	}

	node32 := node23 {
		item: 3,
		left: &node5,
	}
	node2 := node23 {
		item: 2,
		left: &node3,
		right: &node9,
	}

	root := node23 {
		item: 1,
		left: &node32,
		right: &node2,
	}

	rowToMaxVal := map[int]int{0:root.item}
	findMaxRow(&root, rowToMaxVal, 0)

	

	for key, val := range rowToMaxVal {
		fmt.Println("key val", key, val)
	}


	var slice  []int
	slice = append(slice, root.item)
	findMaxRow2(&root, &slice, 0)

	fmt.Println(slice)

}

func findMaxRow(node *node23, rowToMaxVal map[int]int, level int) {
	if node == nil {
		return
	}
	rowToMaxVal[level] = int(math.Max(float64(rowToMaxVal[level]), float64(node.item)))
	findMaxRow(node.left, rowToMaxVal, level + 1)
	findMaxRow(node.right, rowToMaxVal, level + 1)
}

func findMaxRow2(node *node23, slice *[]int, level int) {
		if node == nil {
			return
		}
		// breadth first search strategy - if will satify left(append) and else will satisfy right(update)
		if level > len((*slice)) - 1 {
			(*slice) = append((*slice), node.item)
		} else {
			(*slice)[level] = int(math.Max(float64((*slice)[level]), float64(node.item)))
		}
	
		findMaxRow2(node.left, slice, level + 1)
		findMaxRow2(node.right, slice, level + 1)
	}