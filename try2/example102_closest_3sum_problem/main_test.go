package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThreeSumClosestNumber(t *testing.T) {
	Convey("testing 3 sum", t, func() {
		Convey("success testing", func() {
			data := []struct {
				input     []int
				targetSum int
				output    int
			}{
				{input: []int{-1, 2, 1, -4}, targetSum: 1, output: 2},
				{input: []int{0, 0, 0}, targetSum: 1, output: 0},
				{input: []int{-2, -4, 6, 3, 7}, targetSum: 2, output: 1},
				{input: []int{10, 2, 30, 49, 8}, targetSum: 50, output: 48},
				{input: []int{1, 0, 5, 0, 3}, targetSum: 100, output: 9},
			}

			for _, d := range data {
				So(threeSumClosestNumber(d.input, d.targetSum), ShouldEqual, d.output)
			}
		})
	})
}
