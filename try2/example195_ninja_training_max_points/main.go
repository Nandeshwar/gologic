package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		{10, 20, 30},
		{3, 2, 1},
		{1, 5, 2},
		{1, 3, 2},
	}

	/*
		1st day max activity point = 30
		2nd day max activity point = 3
		3rd day max activiy point = 5
		4th day max activiy point = 2( can not pick 3 because privious day picked same activiy
		                               so next day will be different activity

		ouput should be 30 + 3 + 5 + 2 = 40

	*/

	mp := maxPoints(a, len(a)-1, 3) // 3 means no activity
	fmt.Println("max points=", mp)
}

func maxPoints(a [][]int, day, lastActiviy int) int {
	if day == 0 {
		var m int

		// 3 for activity: there are 3 activities
		for i := 0; i < 3; i++ {
			if i != lastActiviy {
				m = max(m, a[0][i])
			}
		}
		return m
	}

	var m int
	for i := 0; i < 3; i++ {
		if i != lastActiviy {
			points := a[day][i] + maxPoints(a, day-1, i)
			m = max(m, points)
		}
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
