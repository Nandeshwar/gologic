package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinRope(t *testing.T) {
	Convey("testing TestMinRope", t, func() {
		Convey("success testing", func() {
			data := []struct {
				input  []int
				output int
			}{
				{input: []int{2, 5, 4, 8, 6, 9}, output: 85},
			}

			for _, d := range data {
				So(min_rope_weight(d.input), ShouldEqual, d.output)
			}
		})
	})
}
