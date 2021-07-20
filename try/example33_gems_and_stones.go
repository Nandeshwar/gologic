package main

import (
	"fmt"
	"strings"
)

func main() {
	gems := "aA"
	stones := "aANan"

	

	// expectedGemsCount = 3

	 gemsCnt := 0

	 gemsCnt2 := 0

	for _, s := range stones {

		if strings.Index(gems, string(s)) > -1 {
			gemsCnt2++
		}

		for _, g := range gems {
			if g == s {
				gemsCnt++
			}
		}
	}

	fmt.Println(gemsCnt)
	fmt.Println(gemsCnt2)
}