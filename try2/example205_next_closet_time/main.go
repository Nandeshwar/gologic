package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "19:34" // output: "19:39"
	input = "23:59"  // output: "22:22"

	fmt.Println(nextClosestTime2(input))
	fmt.Println("----------------")
	fmt.Println(nextClosestTime(input))
}

func nextClosestTime2(input string) string {
	a := strings.Split(input, ":")
	hhStr := a[0]
	mmStr := a[1]

	hh, _ := strconv.Atoi(hhStr)
	mm, _ := strconv.Atoi(mmStr)

	mm += hh * 60 // when divided by 60 then get hour,  and mod by 60 will give minutes

	fmt.Println("mm=", mm)

	// strore all digits in map
	m := make(map[rune]struct{})
	for _, v := range input {
		if v == ':' {
			continue
		}

		m[v] = struct{}{}
	}

	// keep adding next minute and check if it is valid i.e each digit is inside m
	for {
		mm = (mm + 1) % (24 * 60) // 24 * 60 = 1440, 23:59 = 1439 and if next minute goes more, find mod and proceed
		nextTime := fmt.Sprintf("%02d:%02d", mm/60, mm%60)

		fmt.Println("nextTime=", nextTime)

		isValid := true
		for _, digit := range nextTime {
			if digit == ':' {
				continue
			}
			_, ok := m[digit]
			if !ok {
				isValid = false
			}
		}
		if isValid {
			return nextTime
		}
	}

}

func nextClosestTime(input string) string {
	a := strings.Split(input, ":")
	hhStr := a[0]
	mmStr := a[1]

	hh, _ := strconv.Atoi(hhStr)
	mm, _ := strconv.Atoi(mmStr)

	mm += hh * 60 // when divided by 60 then get hour,  and mod by 60 will give minutes

	fmt.Println("mm=", mm)

	// strore all digits in map
	m := make(map[int]struct{})
	for _, v := range input {
		if v == ':' {
			continue
		}

		m[int(v-'0')] = struct{}{}
	}

	// keep adding next minute and check if it is valid i.e each digit is inside m
	for {
		mm = (mm + 1) % (24 * 60)
		fmt.Println("new mm=", mm) // ex: 1939
		nextTime := []int{
			mm / 60 / 10, // 1
			mm / 60 % 10, // 9
			mm % 60 / 10, // 3
			mm % 60 % 10, // 9
		}

		fmt.Println("nextTime=", nextTime)

		isValid := true
		for _, digit := range nextTime {
			_, ok := m[digit]
			if !ok {
				isValid = false
			}
		}
		if isValid {
			return fmt.Sprintf("%02d:%02d", mm/60, mm%60)
		}
	}

}
