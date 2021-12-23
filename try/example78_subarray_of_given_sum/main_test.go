package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubarrayBySum(t *testing.T) {
	Convey("testing subarrayBySum", t, func() {
		Convey("success testing", func() {
			data := []struct {
				inputArr  []int
				inputSum  int
				outputArr []int
			}{
				{inputArr: []int{1, 2, 3, 4, 5}, inputSum: 7, outputArr: []int{3, 4}},
				{inputArr: []int{1, 2, 3, 4, 5, 11}, inputSum: 16, outputArr: []int{5, 11}},
				{inputArr: []int{1, 2, 3, 4, 5}, inputSum: 1, outputArr: []int{1}},
			}

			for _, d := range data {
				So(subArrayBySum(d.inputArr, d.inputSum), ShouldResemble, d.outputArr)
			}
		})
	})
}
