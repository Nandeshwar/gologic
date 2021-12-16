package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBubbleSort(t *testing.T) {
	Convey("testing Bubbule Sort", t, func() {
		Convey("success testing bubble sort", func() {
			data := []struct {
				input  []int
				output []int
			}{
				{input: []int{5, 2, 1, 7, 8, 3}, output: []int{1, 2, 3, 5, 7, 8}},
			}

			for _, d := range data {
				So(bubbleSort(d.input), ShouldResemble, d.output)
			}
		})
	})
}
