package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFibonacci(t *testing.T) {
	Convey("testing Fibonacci", t, func() {
		Convey("success: testing fibonacci series", func() {
			inputExpectations := []struct {
				input       int
				expectation []int
			}{
				{2, []int{0, 1}},
				{3, []int{0, 1, 1}},
				{4, []int{0, 1, 1, 2}},
			}

			for _, s := range inputExpectations {
				So(fibonacci(s.input, -1, 1, []int{}), ShouldResemble, s.expectation)
			}
		})
	})
}
