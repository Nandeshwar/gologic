package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKnapsackProblem(t *testing.T) {
	Convey("Testing Knapsack Problem", t, func() {
		Convey("success testing knapsack problem", func() {
			data := []struct {
				inputW    []int
				inputV    []int
				bagWeight int
				n         int
				output    int
			}{
				{[]int{2, 5, 1, 3, 4}, []int{15, 14, 10, 45, 30}, 7, 5, 75},
			}

			for _, d := range data {
				So(knapsackProblem(d.inputW, d.inputV, d.bagWeight, d.n), ShouldEqual, d.output)
			}
		})
	})
}
