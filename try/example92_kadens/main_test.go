package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKadensAlgorithm(t *testing.T) {
	Convey("testing kadens Algorithm", t, func() {
		Convey("success testing kadens algorithm", func() {
			data := []struct {
				arr    []int
				output int
			}{
				{arr: []int{1, 0, -1, 2, 4, 9, -7, 8}, output: 16},
			}

			for _, d := range data {
				So(findMaxSubArr(d.arr), ShouldEqual, d.output)
			}
		})
	})
}
