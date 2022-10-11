package main

import (
	"fmt"
	"strconv"
)

func main() {
	// take 1 digit or/and 2 digit
	// result: abc, ic, aw
	printEncoding("123", "")
}

func printEncoding(s, ans string) {
	if len(s) == 0 {
		fmt.Println(ans)
		return
	}

	// ignore 0
	if len(s) == 1 {
		if s == "0" {
			return
		} else {
			singleCh := rune(s[0]) - '0'
			singleChEncodedDigit := 'a' + singleCh - 1
			ans = ans + string(singleChEncodedDigit)
			fmt.Println(ans)
		}
	} else {
		// in case of more than 1 character
		// start with 1 character
		firstCh := s[0]
		restCh := s[1:]

		if firstCh == '0' {
			return
		} else {
			firstChCode := firstCh - '0'
			firstChEncodedDigit := 'a' + firstChCode - 1
			printEncoding(restCh, ans+string(firstChEncodedDigit))
		}

		// start with two characters
		twoCh := s[0:2]
		restCh = s[2:]

		twoDigitInt, _ := strconv.Atoi(twoCh)
		if twoDigitInt <= 26 {

			twoChEncodedDigit := 'a' + twoDigitInt - 1
			printEncoding(restCh, ans+string(twoChEncodedDigit))
		}
	}
}
