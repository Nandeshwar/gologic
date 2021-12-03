package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGCD(t *testing.T) {
	Convey("testing gcd", t, func() {
		Convey("success testing gcd", func() {
			data := []struct {
				inputA int
				inputB int
				output int
			}{
				{24, 40, 8},
			}

			for _, d := range data {
				So(gcd(d.inputA, d.inputB), ShouldEqual, d.output)
			}
		})
	})
}
