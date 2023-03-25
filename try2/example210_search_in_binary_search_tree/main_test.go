package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinaryTreeSearchTree(t *testing.T) {
	Convey("testing binary search tree", t, func() {
		Convey("success: items  found", func() {

			twenty := &Tree{item: 20}
			six := &Tree{item: 6}
			two := &Tree{item: 2}
			five := &Tree{item: 5, left: two, right: six}
			root := &Tree{item: 10, left: five, right: twenty}

			data := []struct {
				input1 *Tree
				input2 int
				output int
			}{
				{input1: root, input2: 5, output: 5},
				{input1: root, input2: 20, output: 20},
			}

			for _, d := range data {
				node := findBinarySearchTree(d.input1, d.input2)
				So(node.item, ShouldEqual, d.output)
			}
		})
	})
}
