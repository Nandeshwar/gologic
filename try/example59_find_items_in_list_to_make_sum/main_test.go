package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPartitionArraySum(t *testing.T) {
	Convey("test Partition Array", t, func() {
		Convey("success: testing parition array sum", func() {
			d := []struct {
				inputArr  []int
				outputArr [][]int
			}{
				{[]int{1, 2, 7, 6, 4}, [][]int{[]int{1, 2, 7}, []int{6, 4}}},
				{[]int{1, 6, 7, 2, 4}, [][]int{[]int{1, 7, 2}, []int{6, 4}}},
			}

			for _, data := range d {
				So(partitionArraySum(data.inputArr), ShouldResemble, data.outputArr)
			}
		})
	})
}
