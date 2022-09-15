package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlusOne(t *testing.T) {
	Convey("Testing Plus One", t, func() {
		Convey("success tesing plus one", func() {
			data := []struct {
				input  []int
				output []int
			}{
				{[]int{1, 2, 9}, []int{1, 3, 0}},
				{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
			}

			for _, d := range data {
				So(plusOne(d.input), ShouldResemble, d.output)
			}
		})
	})
}
