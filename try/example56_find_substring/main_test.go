package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindsubstring(t *testing.T) {
	Convey("test find substring", t, func() {
		Convey("success finding substring abc in abxyabcz", func() {
			actual := findSubstring("abxyabcz", "ab")
			So(actual, ShouldBeTrue)
		})
	})
}
