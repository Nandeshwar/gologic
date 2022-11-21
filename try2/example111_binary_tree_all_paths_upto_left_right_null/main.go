package main

import "fmt"

type Tree struct {
	item  int
	left  *Tree
	right *Tree
}

/*
		10
	  7   9
	5  6


*/
func main() {
	five := Tree{item: 5}
	six := Tree{item: 6}
	t := &Tree{item: 10, left: &Tree{item: 7, left: &five, right: &six}, right: &Tree{item: 9}}

	paths := allPaths(t, []int{})
	fmt.Println(paths)

}

func allPaths(t *Tree, paths []int) [][]int {
	paths = append(paths, t.item)

	if t.left == nil && t.right == nil {
		fmt.Println("paths=", paths)
		return [][]int{paths}
	}

	paths1 := allPaths(t.left, paths)
	paths2 := allPaths(t.right, paths)

	return append(paths1, paths2...)
}

/*
output:
paths= [10 7 5]
paths= [10 7 6]
paths= [10 9]
[[10 7 5] [10 7 6] [10 9]]
*/
