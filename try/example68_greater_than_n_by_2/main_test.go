package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindRepeatedNumberGreaterThanNBy2(t *testing.T) {
	Convey("testing findRepeatedNumberGreaterThanNBy2", t, func() {
		Convey("success testing", func() {
			data := []struct {
				input  []int
				output int
			}{
				{input: []int{2, 1, 1, 3, 1}, output: 1},
				{input: []int{3, 3, 3, 3, 1}, output: 3},
			}

			for _, d := range data {
				So(findRepeatedNumberGreaterThanNBy2(d.input), ShouldEqual, d.output)
			}
		})
	})
}
