package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPathExistsInGrapth(t *testing.T) {
	Convey("testing Path exists in graph", t, func() {
		Convey("suceess testing if path exists", func() {
			data := []struct {
				graphs []EdgesWrapper
				begin  int
				end    int
				result bool
			}{
				{createGraph(), 0, 5, true},
				{createGraph(), 0, 11, false},
			}

			for _, d := range data {
				So(checkPath(d.graphs, d.begin, d.end), ShouldEqual, d.result)
			}
		})
	})
}
