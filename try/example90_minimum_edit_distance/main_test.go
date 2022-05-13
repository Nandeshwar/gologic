package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("testing Main", t, func() {
		Convey("success testing min distance ", func() {
			data := []struct {
				inputStr1 string
				inputStr2 string
				output    int
			}{
				{inputStr1: "abc", inputStr2: "bcd", output: 2},
				{inputStr1: "abc", inputStr2: "", output: 3},
				{inputStr1: "", inputStr2: "bcd", output: 3},
				{inputStr1: "abc", inputStr2: "abe", output: 1},
			}

			for _, d := range data {
				So(min_edit_distance(d.inputStr1, d.inputStr2), ShouldEqual, d.output)
			}
		})
	})
}
