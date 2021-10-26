package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShortestPath(t *testing.T) {
	Convey("Test: shortest path", t, func() {
		Convey("success: testing sortest path", func() {
			testInputOutputData := []struct {
				inputArr           [][]int
				inputX             int
				inputY             int
				expectedPathLength int
			}{
				{[][]int{
					{1, 1},
					{0, 0}}, 0, 1, 1},
				{[][]int{
					{1, 1},
					{0, 1},
					{1, 1, 1}}, 2, 2, 4},
			}

			for _, data := range testInputOutputData {
				So(findShortestPath(data.inputArr, data.inputX, data.inputY), ShouldEqual, data.expectedPathLength)
			}
		})
	})
}
