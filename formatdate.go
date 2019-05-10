package main

import "time"
import "fmt"

func main() {
	past := time.Now().AddDate(0, 0, -7).UTC()
	pastUTC := time.Date(past.Year(), past.Month(), past.Day(), 0, 0, 0, 0, time.UTC)
	pastUTCStr := pastUTC.Format("2006-01-02T15:04:00Z")
	fmt.Println(pastUTCStr)
}
