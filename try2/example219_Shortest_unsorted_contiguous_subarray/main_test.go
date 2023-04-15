package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShortestUnsortedArray(t *testing.T) {
	Convey("testing function ShortestUnsortedSubArray", t, func() {
		Convey("success testing", func() {
			data := []struct {
				input  []int
				output []int
			}{
				{input: []int{2, 6, 4, 8, 10, 9, 15}, output: []int{6, 4, 8, 10, 9}},
				{input: []int{1, 2, 3, 4}, output: []int{}},
			}

			for _, d := range data {
				So(findShortestUnsortedSubArray(d.input), ShouldResemble, d.output)
			}
		})
	})
}
