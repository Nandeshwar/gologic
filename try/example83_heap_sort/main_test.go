package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHeapSort(t *testing.T) {
	Convey("testing heap sort", t, func() {
		Convey("success testing heap sort", func() {
			data := []struct {
				input  []int
				output []int
			}{
				{input: []int{2, 1, 5, 4, 3}, output: []int{1, 2, 3, 4, 5}},
			}
			
			for _, d := range data {
				heapSort(d.input)
				So(d.input, ShouldResemble, d.output)
			}
		})
	})
}
