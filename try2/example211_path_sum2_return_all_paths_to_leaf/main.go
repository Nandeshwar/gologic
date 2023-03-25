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
				10
			5	    8
			       3
		  2   6


		output: path for 21
		[[10 5 6] [10 8 3]]
	*/

	three := &Tree{item: 3}
	eight := &Tree{item: 8, left: three}
	six := &Tree{item: 6}
	two := &Tree{item: 2}
	five := &Tree{item: 5, left: two, right: six}
	root := &Tree{item: 10, left: five, right: eight}

	paths := [][]int{}
	currentPath := []int{}
	sum := 21
	allPaths(root, &paths, currentPath, sum)

	fmt.Println(paths)
}

func allPaths(root *Tree, paths *[][]int, currentPath []int, sum int) {
	if root == nil {
		return
	}

	currentPath = append(currentPath, root.item)
	if root.item == sum && root.left == nil && root.right == nil {
		*paths = append(*paths, currentPath)
		return
	}

	allPaths(root.left, paths, currentPath, sum-root.item)
	allPaths(root.right, paths, currentPath, sum-root.item)
}
