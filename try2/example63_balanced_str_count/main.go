package main

import (
	"fmt"
)

func main() {
	fmt.Println(balancedStrCount("RRLLLLLRRR")) // expected output 2 here. RRLL, LLLRRR
}

func balancedStrCount(str string) int {
	totalBalanceString := 0
	cnt := 0
	for _, v := range str {
		s := string(v)

		switch s {
		case "R", "r":
			cnt++
		case "L", "l":
			cnt--
		}

		if cnt == 0 {
			totalBalanceString++
		}
	}
	return totalBalanceString
}
