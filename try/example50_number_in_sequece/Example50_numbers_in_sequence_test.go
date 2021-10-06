package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"reflect"
)



func TestNewSequence(t *testing.T) {
	Convey("Test newSequence Test", t, func(){
		Convey("success: preparing new sequece", func() {
			expectedList := []int{5, 10, 30,  20,  11}
			numList := []int{10, 30, 5, 20, 11}
			opList := []string{"<", "<", ">", ">"}

			actual := newSequece(numList, opList)
		
			So(actual, ShouldResemble, expectedList)
			So(reflect.DeepEqual(actual, expectedList), ShouldBeTrue)
		})
	})
}