package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRainwaterTrapping(t *testing.T) {
	Convey("testing rainwater trap", t, func() {
		Convey("success testing", func() {
			data := []struct {
				inputBuildingHeight []int
				outputTrappedWater  int
			}{
				{[]int{1, 2, 1, 2, 1}, 1},
			}

			for _, d := range data {
				So(getRainWaterTrappedBlock(d.inputBuildingHeight), ShouldEqual, d.outputTrappedWater)
			}
		})
	})
}
