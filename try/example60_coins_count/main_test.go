package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCoinNumberToMakeSum(t *testing.T) {
	Convey("Test conin numbers to make sum", t, func() {
		Convey("success: testing coin numbers to make sum", func() {
			data := []struct {
				inputCoinList   []int
				inputSum        int
				outputCoinCount int
			}{
				{[]int{1, 7, 5}, 18, 4},
				{[]int{1, 7, 5, 2}, 18, 4},
				{[]int{1, 8, 2}, 18, 3},
			}

			for _, d := range data {

				So(coinCountToMakeSum(d.inputCoinList, d.inputSum), ShouldEqual, d.outputCoinCount)
			}
		})
	})
}
