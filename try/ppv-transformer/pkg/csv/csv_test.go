package csv

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCsv(t *testing.T) {
	Convey("Testing CSV", t, func() {
		Convey("success: parsing 2 line csv. 52th field should be replaced with 53th field", func() {
			ppvList, err := processCSV("PPV_SMALL_2line.TXT")
			So(err, ShouldBeNil)
			So(len(ppvList), ShouldEqual, 2)
			So(ppvList[0].Fields81Output[51], ShouldEqual, ppvList[0].Fields72[52])
			So(ppvList[0].Fields81Output[51], ShouldEqual, "")
			So(ppvList[0].Fields81Output[55], ShouldEqual, "TVMA")
		})
		Convey("failure: 52th field empty for 72fields CSV", func() {
			_, err := processCSV("PPV_SMALL_52th_field_empty.TXT")
			So(err.Error(), ShouldEqual, "52th field is empty at line=1")
		})
		Convey("failure: if file does not exist", func() {
			err := TransformCSV("abc", "", "", "", "")
			So(err.Error(), ShouldContainSubstring, "csv file=abc does not exist. error")
		})
	})
}
