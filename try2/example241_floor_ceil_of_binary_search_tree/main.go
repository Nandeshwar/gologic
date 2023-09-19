package main

/*
Binary Search tree insertion visualization: https://www.cs.usfca.edu/~galles/visualization/BST.html
Inorder traversal of binary search tree is always in ascending order
*/

import (
	"fmt"
	"math"
)

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

func main() {
	/*
	   			10
	   	  5           20
	   2           19   21
	  1  4

	*/

	twentyOne := &Tree{item: 21}
	nineteen := &Tree{item: 19}
	four := &Tree{item: 4}
	two := &Tree{item: 2, right: four, left: &Tree{item: 1}}

	t := new(Tree)
	t.item = 10
	t.left = &Tree{item: 5, left: two}
	t.right = &Tree{item: 20, left: nineteen, right: twentyOne}

	f, c := floorCeil(t, 5)
	fmt.Println("floor of 5=", f)
	fmt.Println("ceil of 5", c)

	fmt.Println("key = 3: Algorithm2 floor=", floorAlgo2(t, 3))
	fmt.Println("key = 3: Algorithm2 ceil=", ceilAlgo2(t, 3))

}

func floorCeil(t *Tree, item int) (int, int) {
	f := 0
	c := 0
	result := []int{}
	inorder(t, &result)

	fmt.Println("result=", result)
	for i := 0; i < len(result); i++ {
		if item == result[i] {
			f = item
			c = item
			break
		}
		if result[i] > item {
			f = result[i-1]
			c = result[i]
			break
		}
	}

	return f, c
}

func inorder(t *Tree, result *[]int) {
	if t == nil {
		return
	}

	inorder(t.left, result)
	*result = append(*result, t.item)
	inorder(t.right, result)
}

func floorAlgo2(t *Tree, item int) int {
	f := math.MaxInt

	for t != nil {
		if t.item == item {
			return t.item
		}

		// want to get smaller value, so if item in tree > key , that means smaller value in left side
		if t.item > item {
			t = t.left
		} else if t.item < item {
			f = t.item
			t = t.right
		}
	}

	return f
}

func ceilAlgo2(t *Tree, item int) int {
	f := math.MinInt

	for t != nil {
		if t.item == item {
			return t.item
		}

		if t.item < item {
			t = t.right
		} else if t.item > item {
			f = t.item
			t = t.left
		}
	}

	return f
}
