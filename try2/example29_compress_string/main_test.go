package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("testMain", t, func() {
		Convey("success", func() {
			data := []struct {
				input  string
				output string
			}{
				{input: "aabbccc", output: "a2b2c3"},
				{input: "abbb", output: "ab3"},
			}

			for _, d := range data {
				result := compressStr(d.input)
				So(result, ShouldEqual, d.output)
			}
		})
	})
}
