package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsMountain(t *testing.T) {
	Convey("testing", t, func() {
		Convey("success testing", func() {
			data := []struct {
				a []int
				b bool
			}{
				{[]int{1, 2, 3}, false},
				{[]int{1, 2, 3, 3}, false},
				{[]int{1, 2, 3, 2, 1}, true},
			}

			for _, d := range data {
				So(isMountain(d.a), ShouldEqual, d.b)
			}
		})
	})
}
