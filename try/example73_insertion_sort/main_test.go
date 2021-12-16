package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertionSort(t *testing.T) {
	Convey("testing insertion sort", t, func() {
		Convey("success testing insertion sort", func() {
			data := []struct {
				input  []int
				output []int
			}{
				{input: []int{8, 1, 7, 2, 5}, output: []int{1, 2, 5, 7, 8}},
			}

			for _, d := range data {
				insertionSort(d.input)
				So(d.input, ShouldResemble, d.output)
			}
		})
	})
}
