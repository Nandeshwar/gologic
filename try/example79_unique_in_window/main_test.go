package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniqueItemsInEachWindow(t *testing.T) {
	Convey("Test unique items in each window", t, func() {
		Convey("success testing", func() {
			data := []struct {
				inputArr    []int
				inputWindow int
				outputArr   []int
			}{
				{inputArr: []int{1, 2, 3, 3, 4, 4}, inputWindow: 4, outputArr: []int{3, 3, 2}},
			}

			for _, d := range data {
				So(uniqueInWindow(d.inputArr, d.inputWindow), ShouldResemble, d.outputArr)
			}
		})
	})
}
