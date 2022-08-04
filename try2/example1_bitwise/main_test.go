package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSwap(t *testing.T) {
	Convey("testing swap", t, func() {
		Convey("success testing swap", func() {
			data := []struct {
				input1  int
				input2  int
				output1 int
				output2 int
			}{
				{1, 2, 2, 1},
				{3, 4, 4, 3},
			}

			for _, d := range data {
				a, b := swap(d.input1, d.input2)
				So(a, ShouldEqual, d.output1)
				So(b, ShouldEqual, d.output2)
			}
		})
	})
}


