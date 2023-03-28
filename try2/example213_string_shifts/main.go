package main

import (
	"fmt"
)

func main() {
	s := "abcdefg"
	// 0 - left shift
	// 1 - right shift

	//output: efgabcd

	shift := [][]int{
		{1, 1},
		{1, 1},
		{0, 2},
		{1, 3},
	}

	// Algorithm
	// 1. collect left shift count
	// 2. collect right shift count
	// 3. see which is greater,that decides if left or right shift
	// if left > right : shift count will (left - right) % len(s)
	// if right > left : shift count will be (right - left) % len(s)

	fmt.Println(shiftString(s, shift))
}

func shiftString(s string, shift [][]int) string {
	var result string
	left := 0
	right := 0
	for _, a := range shift {
		switch a[0] {
		case 0:
			left += a[1]
		case 1:
			right += a[1]
		}
	}

	shiftDirection := "left"
	shiftCount := left

	if right > left {
		shiftDirection = "right"
		shiftCount = right - left
	}

	if right < left {
		shiftCount = left - right
	}

	shiftCount = shiftCount % len(s)

	fmt.Println("shiftDirection=", shiftDirection)
	fmt.Println("shiftCount=", shiftCount)

	if shiftDirection == "right" {
		index := len(s) - shiftCount
		rightPortion := s[index:]
		result = rightPortion + s[0:index]
	} else {
		index := shiftCount
		leftPortion := s[0:index]
		result = s[index:] + leftPortion
	}

	return result
}
