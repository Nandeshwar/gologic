/*
    To see bytes used by per operation
	go test -benchmem -run=^$ -bench .
*/

package main

import (
	"testing"
)

func BenchmarkDeleteAllNodes(b *testing.B) {
	/*
			   1
		     2   3
		    4 5    7
		            8

	*/

	node8 := &Node{item: 8}
	node7 := &Node{item: 7, right: node8}

	node5 := &Node{item: 5}
	node4 := &Node{item: 4}

	node2 := &Node{item: 2, left: node4, right: node5}
	node3 := &Node{item: 3, right: node7}
	start := &Node{item: 1, left: node2, right: node3}
	for n := 0; n < b.N; n++ {

		deleteAllNodes(start)
	}
}

func BenchmarkDeleteAllNodesRecursive(b *testing.B) {
	/*
			   1
		     2   3
		    4 5    7
		            8

	*/

	node8 := &Node{item: 8}
	node7 := &Node{item: 7, right: node8}

	node5 := &Node{item: 5}
	node4 := &Node{item: 4}

	node2 := &Node{item: 2, left: node4, right: node5}
	node3 := &Node{item: 3, right: node7}
	start := &Node{item: 1, left: node2, right: node3}

	for n := 0; n < b.N; n++ {

		deleteAllNodesRecursive(start)
	}
}
