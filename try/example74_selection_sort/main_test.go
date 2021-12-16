package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSelectionSort(t *testing.T) {
	Convey("test selection sort", t, func() {
		Convey("success testing selection sort", func() {
			data := []struct {
				input  []int
				output []int
			}{
				{input: []int{8, 1, 5, 7, 2}, output: []int{1, 2, 5, 7, 8}},
			}

			for _, d := range data {
				selectionSort(d.input)
				So(d.input, ShouldResemble, d.output)
			}
		})
	})
}
