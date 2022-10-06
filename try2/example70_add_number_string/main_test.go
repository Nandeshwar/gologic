package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBigNumberString(t *testing.T) {
	Convey("Testing BigNumberAddString", t, func() {
		Convey("Test addition of two numbers", func() {
			data := []struct {
				num1   string
				num2   string
				result string
			}{
				{num1: "10", num2: "20", result: "30"},
				{num1: "99", num2: "99", result: "198"},
				{num1: "1111111111111111111111111", num2: "1111111111111111111111111", result: "2222222222222222222222222"},
			}

			for _, v := range data {
				So(add(v.num1, v.num2), ShouldEqual, v.result)
			}
		})
	})
}
