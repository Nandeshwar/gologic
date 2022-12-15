package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPreviousSmallestNumber(t *testing.T) {
	Convey("testing Previous Smallest number", t, func() {
		Convey("success testing", func() {
			data := []struct {
				input  []int
				output []int
			}{
				{input: []int{4, 10, 8, 15, 3}, output: []int{-1, 4, 4, 8, -1}},
				{input: []int{1, 2, 3}, output: []int{-1, 1, 2}},
				{input: []int{6, 5, 4}, output: []int{-1, -1, -1}},
			}

			for _, d := range data {
				So(getPreviousSmallestElement(d.input), ShouldResemble, d.output)
			}
		})
	})
}
