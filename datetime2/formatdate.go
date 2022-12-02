package main

import (
	"fmt"
	"time"
)

func main() {

	localTime := time.Now().Local()
	utcTime := time.Now().UTC()
	defaultTimeNow := time.Now()

	fmt.Println("localTime", localTime)
	fmt.Println("UTC time=", utcTime)
	fmt.Println("default timeNow", defaultTimeNow)

	past := time.Now().AddDate(0, 0, -7).UTC()

	pastUTC := time.Date(past.Year(), past.Month(), past.Day(), 0, 0, 0, 0, time.UTC)
	pastUTCStr := pastUTC.Format("2006-01-02T15:04:00Z")
	fmt.Println(pastUTCStr)
	fmt.Println(past.Format("2006-01-02T15:04:00Z"))
	past = past.Add(1 * time.Hour)
	fmt.Println(past.Format("2006-01-02T15:04:00Z"))

	hr := 1
	past = past.Add(time.Duration(hr) * time.Hour)
	fmt.Println(past.Format("2006-01-02T15:04:00Z"))

	/*
					output:
					bash-3.2$ go run formatdate.go
		localTime 2022-12-02 15:01:18.417084 -0700 MST
		UTC time= 2022-12-02 22:01:18.417084 +0000 UTC
		default timeNow 2022-12-02 15:01:18.417084 -0700 MST m=+0.000085589
		2022-11-25T00:00:00Z
		2022-11-25T22:01:00Z
		2022-11-25T23:01:00Z
		2022-11-26T00:01:00Z

	*/
}

// 1. No Date should be in past
// 2. Reservation date should not be after arrival date
// 3. checkout date should not be before arrival date or reservationDate
func validateDates(reservationDateStr, arrivalDateStr, checkoutDateStr string) bool {
	reservationDate, err := strToDate(reservationDateStr)
	if err != nil {
		fmt.Println("error parsing reservation date", err)
		return false
	}

	arrivalDate, err := strToDate(arrivalDateStr)
	if err != nil {
		fmt.Println("error parsing arrival date", err)
		return false
	}

	checkoutDate, err := strToDate(checkoutDateStr)
	if err != nil {
		fmt.Println("error parsing checkout date", err)
		return false
	}

	today := time.Now()
	// set time to zero both lines below work
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, &time.Location{})
	//today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	if reservationDate.Before(today) {
		fmt.Printf("\n reservation date=%v should not be in past. todayDate=%v", reservationDate, today)
		return false
	}

	if arrivalDate.Before(today) {
		fmt.Printf("\n arrival  date=%v should not be in past. todayDate=%v", arrivalDate, today)
		return false
	}

	if checkoutDate.Before(today) {
		fmt.Printf("\ncheckout date=%v date should not be in past. todayDate=%v", checkoutDate, today)
		return false
	}

	return true
}

func strToDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
