package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThreeSum(t *testing.T) {
	Convey("test three sum", t, func() {
		Convey("success:", func() {
			arr := []int{-1, 0, 1, 2, -1, -4}
			expectedList := [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			}
			result := ThreeSum(arr)
			So(result, ShouldResemble, expectedList)
		})
	})
}
