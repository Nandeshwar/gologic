package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestMain(t *testing.T) {
	Convey("Testing Main function", t, func() {
		Convey("Success main: ", func() {
			So(1, ShouldEqual, 1)
		})

		str := "pwwwkew"
		expectedLen := 3
		Convey(fmt.Sprintf("success: substring length=%d in string =%s  without repetion of characters.", expectedLen, str), func(){
			actual := findLongestLenOfSubstring_withoutrRepetition(str)
			So(actual, ShouldEqual, expectedLen)
		})

		str = "abccdef"
		expectedLen = 4
		Convey(fmt.Sprintf("success: substring length=%d in string =%s  without repetion of characters.", expectedLen, str), func(){
			actual := findLongestLenOfSubstring_withoutrRepetition(str)
			So(actual, ShouldEqual, expectedLen)
		})
	})
}