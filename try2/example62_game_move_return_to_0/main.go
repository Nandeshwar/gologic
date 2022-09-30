package main

import (
	"fmt"
)

func main() {
	fmt.Println(moveReturnTo0("RLUD"))
}

func moveReturnTo0(move string) bool {
	rCnt := 0
	uCnt := 0

	for _, v := range move {
		switch string(v) {
		case "R", "r":
			rCnt++
		case "L", "l":
			rCnt--
		case "U", "u":
			uCnt++
		case "D", "d":
			uCnt--
		}
	}

	return rCnt == 0 && uCnt == 0
}
