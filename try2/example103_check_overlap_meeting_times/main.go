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
		{8, 11},
		{10, 11},
		{9, 11},
		{11, 12},
	}

	fmt.Println(isOverlap(meetingTimesList))

	fmt.Println(allOverlap(meetingTimesList))
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

func allOverlap(meetingTimesList []MeetingTimes) []MeetingTimes {
	var overlapMeetings []MeetingTimes

	sort.Slice(meetingTimesList, func(i, j int) bool {
		return meetingTimesList[i].start < meetingTimesList[j].start
	})

	goodMeeting := meetingTimesList[0]
	for i := 1; i < len(meetingTimesList); i++ {
		if meetingTimesList[i].start < goodMeeting.end {
			overlapMeetings = append(overlapMeetings, meetingTimesList[i])
		} else {
			goodMeeting = meetingTimesList[i]
		}
	}
	return overlapMeetings
}
