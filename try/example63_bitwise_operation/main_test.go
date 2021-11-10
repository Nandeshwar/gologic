package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBitwiseOperator(t *testing.T) {
	Convey("Test bitwise operations", t, func() {
		Convey("success: testing << operator which means multiply by 2", func() {
			data := []struct {
				inputNum int
				result   int
			}{
				{inputNum: 5, result: 10},
				{inputNum: 10, result: 20},
				{inputNum: 9, result: 18},
			}
			for _, d := range data {

				So(multiplyBy2(d.inputNum), ShouldEqual, d.result)
			}
		})

		Convey("success: testing >> operator which means devide by 2", func() {
			data := []struct {
				inputNum int
				result   int
			}{
				{inputNum: 5, result: 2},
				{inputNum: 10, result: 5},
				{inputNum: 9, result: 4},
			}
			for _, d := range data {

				So(divideBy2(d.inputNum), ShouldEqual, d.result)
			}
		})

		Convey("success: testing &1 operator which means return 0 for even and returns 1 for odd", func() {
			data := []struct {
				inputNum int
				result   bool
			}{
				{inputNum: 5, result: false},
				{inputNum: 10, result: true},
				{inputNum: 9, result: false},
			}
			for _, d := range data {

				So(isEven(d.inputNum), ShouldEqual, d.result)
			}
		})

		Convey("success: testing xor operator which is 1 every both input are different", func() {
			data := []struct {
				input1  int
				input2  int
				output1 int
				output2 int
			}{
				{input1: 5, input2: 10, output1: 10, output2: 5},
			}
			for _, d := range data {

				result1, result2 := swapNumber(d.input1, d.input2)
				So(result1, ShouldEqual, d.output1)
				So(result2, ShouldEqual, d.output2)
			}
		})
	})
}
