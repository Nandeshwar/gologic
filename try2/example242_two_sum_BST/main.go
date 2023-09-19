package main

import "fmt"

/*
 check if the tree contains 2 values which sum equals to target value
approach 1:
    inorder traversal, and collect all value in array and that will be sorted order, then use 2 sum technique
approach 2:
    traverse to tree in any order and use map store target - item value, and check if that already present
     this similar to 2 sum technique using space
*/

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

	m := make(map[int]struct{})
	fmt.Println(bstTwoSum(t, 21, m))
	fmt.Println(bstTwoSum(t, 22, m))
	fmt.Println(bstTwoSum(t, 13, m))
}

func bstTwoSum(t *Tree, target int, m map[int]struct{}) bool {
	if t == nil {
		return false
	}

	if bstTwoSum(t.left, target, m) == true {
		return true
	}
	if _, ok := m[target-t.item]; ok {
		return true
	}
	m[t.item] = struct{}{}
	return bstTwoSum(t.right, target, m)
}
