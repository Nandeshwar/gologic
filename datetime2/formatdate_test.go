/*
  How to setup and run test in local
  1.  install goconvey in local
     sudo  go get github.com/smartystreets/goconvey
  2.  Run test
     a. sudo goconvey 
	   or
     b. go test -v 
	    or
	 c. go test 
	     or 
     d. go test ./...
 
*/
package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestValidateDates(t *testing.T) {
	Convey("validate test", t, func() {
		
		pastReservationDate := "2021-09-28"
		Convey("success: validate reservation date, arrival date, checkoutDate", func() {
				reservationDate := "2021-09-29"
				arrivalDate := "2021-09-29"
				checkoutDate := "2021-09-29"
				actual := validateDates(reservationDate, arrivalDate, checkoutDate)
				So(actual, ShouldBeTrue)
			})
			Convey(fmt.Sprintf("failure: reservationDate=%v should not be in the past", pastReservationDate), func() {
				arrivalDate := "2021-09-29"
				checkoutDate := "2021-09-29"
				actual := validateDates(pastReservationDate, arrivalDate, checkoutDate)
				So(actual, ShouldBeFalse)
			})
		})
}