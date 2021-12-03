package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKadensAlgorithm(t *testing.T) {
	Convey("testing kaden's algorith", t, func() {
		Convey("success testing kaden's algorithm", func() {
			data := []struct {
				inputArr  []int
				outputMax int
				outputArr []int
			}{
				{inputArr: []int{-2, -3, 4, -1, -2, 1, 5, -3}, outputMax: 7, outputArr: []int{4, -1, -2, 1, 5}},
			}

			for _, d := range data {
				maxSum, subArr := findMaxValueAndContinousSubarry(d.inputArr)

				fmt.Println("maxSum in test", maxSum)

				So(maxSum, ShouldEqual, d.outputMax)
				So(subArr, ShouldResemble, d.outputArr)

			}
		})
	})
}
