package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPow(t *testing.T) {
	Convey("Test Power of n", t, func() {
		b := 2
		p := 4
		expected := 16
		Convey(fmt.Sprintf("success: pow(%d,%d) is %d", b, p, expected), func() {
			actual := pow(b, p)
			So(actual, ShouldEqual, expected)
		})

	})
}
