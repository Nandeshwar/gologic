package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearchInSortedCircularArry(t *testing.T) {
	Convey("testing Search in Sorted Circular Array", t, func() {
		Convey("success testing", func() {
			data := []struct {
				inputArr    []int
				searchItem  int
				outputIndex int
			}{
				{inputArr: []int{7, 1, 2, 3, 4}, searchItem: 2, outputIndex: 2},
			}

			for _, d := range data {
				So(searchInSortedCircularArr(d.inputArr, d.searchItem), ShouldEqual, d.outputIndex)
			}
		})
	})
}
