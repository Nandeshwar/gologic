package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinBoatRequired(t *testing.T) {
	Convey("testing func minBoatRequired", t, func() {
		Convey("success testing", func() {
			data := []struct {
				input1 []int
				input2 int
				output int
			}{
				{[]int{1, 2}, 3, 1},
				{[]int{3, 5, 3, 4}, 5, 4},
			}

			for _, d := range data {
				So(minBoatRequired(d.input1, d.input2), ShouldEqual, d.output)
			}
		})
	})
}
