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

	nine := &Tree{item: 9}
	eight := &Tree{item: 8}
	four := &Tree{item: 4}
	three := &Tree{item: 3}
	five := &Tree{item: 5, left: three, right: four}
	t1 := &Tree{item: 10, left: five, right: eight}

	t2 := &Tree{item: 10, left: five, right: eight}
	t3 := &Tree{item: 10, left: five, right: nine}

	fmt.Println(isTreeSame(t1, t2))
	fmt.Println(isTreeSame(t1, t3))
}

func isTreeSame(t1, t2 *Tree) bool {
	if t1 == nil || t2 == nil {
		return t1 == nil && t2 == nil
	}

	return t1.item == t2.item && isTreeSame(t1.left, t2.left) && isTreeSame(t1.right, t2.right)
}
