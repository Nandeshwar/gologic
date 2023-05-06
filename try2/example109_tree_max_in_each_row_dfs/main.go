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

  map[1:10 2:8 3:9 4:33]
*/
func main() {
	nine := &Node{item: 9, left: &Node{item: 33}}
	seven := &Node{item: 7}
	eight := &Node{item: 8, right: nine}

	t := &Node{item: 10, left: seven, right: eight}

	m := map[int]int{1: t.item}
	sumInEachRow(t, 1, m)
	fmt.Println(m)

}

func sumInEachRow(t *Node, row int, m map[int]int) {
	if t == nil {
		return
	}

	v := m[row]
	m[row] = max(v, t.item)

	sumInEachRow(t.left, row+1, m)
	sumInEachRow(t.right, row+1, m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
