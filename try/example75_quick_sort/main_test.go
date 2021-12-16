package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuickSort(t *testing.T) {
	Convey("testing quick sort", t, func() {
		Convey("success testing quick sort", func() {
			data := []struct {
				input  []int
				output []int
			}{
				{input: []int{4, 1, 5, 2, 3}, output: []int{1, 2, 3, 4, 5}},
				{input: []int{4, 1, 5, 3, 2}, output: []int{1, 2, 3, 4, 5}},
			}

			for _, d := range data {
				quickSort(d.input)
				So(d.input, ShouldResemble, d.output)
			}
		})
	})
}
