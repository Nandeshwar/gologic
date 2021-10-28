package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJosephProblem(t *testing.T) {
	Convey("test Joseph problem", t, func() {
		Convey("success testing joseph problem", func() {
			data := []struct {
				inputN int
				inputK int
				output int
			}{
				{inputN: 5, inputK: 3, output: 3},
				{inputN: 7, inputK: 4, output: 1},
			}

			for _, d := range data {
				So(JosephProblem(d.inputN, d.inputK), ShouldEqual, d.output)
			}
		})
	})
}
