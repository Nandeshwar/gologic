package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFibonacci(t *testing.T) {
	Convey("testing fibonacci", t, func() {
		Convey("success testing fibonacci", func() {
			data := []struct {
				n int
				f int
			}{
				{n: 5, f: 5},
				{n: 12, f: 144},
			}

			for _, d := range data {
				m := make(map[int]int)
				So(fibonacci(d.n, m), ShouldEqual, d.f)
			}
		})
	})
}
