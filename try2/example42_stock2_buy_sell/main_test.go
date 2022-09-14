package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStockSell(t *testing.T) {
	Convey("testingStockSell", t, func() {
		Convey("success testing algo1", func() {
			data := []struct {
				input  []int
				output int
			}{
				{input: []int{1, 2, 3}, output: 2},
				{input: []int{8, 9, 6, 6, 7}, output: 2},
				{input: []int{8, 9, 6, 6, 7, 1, 3}, output: 4},
			}

			for _, d := range data {
				So(stock2BuySell(d.input), ShouldEqual, d.output)
				So(stock2BuySellAlgo2(d.input), ShouldEqual, d.output)
			}
		})
	})
}
