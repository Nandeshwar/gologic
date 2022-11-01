package main

import (
	"fmt"
	"sort"
)

type MeetingTimes struct {
	start int
	end   int
}

func main() {
	meetingTimesList := []MeetingTimes{
		{3, 6},
		{1, 2},
		{5, 7}, // overlap here: 3-6, then 5-7
		{8, 9},
	}

	fmt.Println(isOverlap(meetingTimesList))
}

func isOverlap(meetingTimesList []MeetingTimes) bool {
	sort.Slice(meetingTimesList, func(i, j int) bool {
		return meetingTimesList[i].start < meetingTimesList[j].start
	})

	for i := 0; i < len(meetingTimesList)-1; i++ {
		if meetingTimesList[i+1].start < meetingTimesList[i].end {
			return true
		}
	}
	return false
}
