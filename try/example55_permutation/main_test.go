package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPermutation(t *testing.T) {
	Convey("test permutation functions", t, func() {
		Convey("success: testing permuations for abc", func() {
			str := "abc"
			left := 0
			right := 2

			expectedStrList := []string{"abc", "acb", "bac", "bca", "cba", "cab"}

			actual := permutation(str, left, right, []string{})
			So(actual, ShouldResemble, expectedStrList)
		})
	})
}
