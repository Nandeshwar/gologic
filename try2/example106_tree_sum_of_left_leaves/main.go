package main

import (
	"fmt"
)

type Node struct {
	item  int
	left  *Node
	right *Node
}

/*
     10
   7   8
         9
        33

  left leaves : 7 and 33 . so sum will be 40
*/
func main() {
	nine := &Node{item: 9, left: &Node{item: 33}}
	seven := &Node{item: 7}
	eight := &Node{item: 8, right: nine}

	t := &Node{item: 10, left: seven, right: eight}

	fmt.Println(leftLeavesSum2(t, false))

}

func leftLeavesSum2(t *Node, left bool) int {

	if t == nil {
		return 0
	}

	if left {
		if t.left == nil && t.right == nil {
			return t.item
		}
	}

	sum1 := leftLeavesSum2(t.left, true)
	sum2 := leftLeavesSum2(t.right, false)
	return sum1 + sum2
}
