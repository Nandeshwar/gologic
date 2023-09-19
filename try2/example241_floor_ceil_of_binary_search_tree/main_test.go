package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFloorCeil(t *testing.T) {

	Convey("testing FloorCeil", t, func() {
		/*
		   			10
		   	  5           20
		   2           19   21
		   	  4
		   	1
		*/

		twentyOne := &Tree{item: 21}
		nineteen := &Tree{item: 19}
		four := &Tree{item: 4}
		two := &Tree{item: 2, right: four, left: &Tree{item: 1}}

		t := new(Tree)
		t.item = 10
		t.left = &Tree{item: 5, left: two}
		t.right = &Tree{item: 20, left: nineteen, right: twentyOne}

		Convey("success testing different combination", func() {
			data := []struct {
				input int
				floor int
				ceil  int
			}{
				{input: 5, floor: 5, ceil: 5},
				{input: 3, floor: 2, ceil: 4},
			}

			for _, d := range data {
				f, c := floorCeil(t, d.input)
				So(f, ShouldEqual, d.floor)
				So(c, ShouldEqual, d.ceil)
			}
		})
	})
}
