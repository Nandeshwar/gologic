package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindFruitsNumberInBasket(t *testing.T) {
	Convey("Test FindFruitsNumberInBasket", t, func() {
		expectedSize := 5
		trees := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
		Convey(fmt.Sprintf("success: max number of fruites=%d in trees=%v", expectedSize, trees), func() {
			actual := findFruitsNumberInBasket(trees)
			So(actual, ShouldEqual, expectedSize)
		})

		expectedSize = 4
		trees = []int{3, 3, 3, 1, 2, 1, 1, 3, 3, 4}
		Convey(fmt.Sprintf("success: max number of fruites=%d in trees=%v", expectedSize, trees), func() {
			actual := findFruitsNumberInBasket(trees)
			So(actual, ShouldEqual, expectedSize)
		})
	})
}
