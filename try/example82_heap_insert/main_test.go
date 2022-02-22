package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHeapInsert(t *testing.T) {
	Convey("testing Heap", t, func() {
		Convey("success testing insert index", func() {
			data := []struct {
				inputArr  []int
				inputItem int
				outputArr []int
			}{
				{inputArr: []int{20, 18, 19, 15, 17}, inputItem: 25, outputArr: []int{0, 25, 18, 20, 15, 17, 19}},
			}

			for _, d := range data {
				create_heap(&d.inputArr)
				heap_insert(&d.inputArr, d.inputItem)
				So(d.inputArr, ShouldResemble, d.outputArr)
			}
		})

		Convey("success testing delete index", func() {
			data := []struct {
				inputArr  []int
				inputItem int
				outputArr []int
			}{
				{inputArr: []int{20, 18, 19, 15, 17}, outputArr: []int{0, 19, 18, 17, 15, 0}},
			}

			for _, d := range data {
				create_heap(&d.inputArr)
				heap_delete(&d.inputArr)
				So(d.inputArr, ShouldResemble, d.outputArr)
			}
		})
	})
}
