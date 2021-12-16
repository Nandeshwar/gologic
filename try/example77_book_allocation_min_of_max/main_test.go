package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinOfMaxBookAllocation(t *testing.T) {
	Convey("testing MinOfMaxBookAllocation", t, func() {
		Convey("success testing minOfMaxBookAllocation", func() {
			data := []struct {
				inputArr            []int
				inputStudent        int
				outputMinOfMaxPages int
			}{
				{inputArr: []int{10, 20, 30, 40}, inputStudent: 2, outputMinOfMaxPages: 60},
			}

			for _, d := range data {
				So(minOfMaxBookAllocation(d.inputArr, d.inputStudent), ShouldEqual, d.outputMinOfMaxPages)
			}
		})
	})
}
