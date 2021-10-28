package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountAllPossbilePathsInMatrix(t *testing.T) {
	Convey("test count all possbile paths in matrix", t, func() {
		Convey("success testing all possible paths", func() {
			data := []struct {
				inputArr       [][]int
				expectedOutput int
			}{
				{[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				}, 6},
			}

			for _, d := range data {
				So(findAllPossiblePathsCountInMatrix(d.inputArr), ShouldEqual, d.expectedOutput)
			}
		})
	})
}
