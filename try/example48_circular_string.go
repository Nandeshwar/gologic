// And Easy solution will be check if (s1 + s1).contains(s2)
package main

import (
	"fmt"
)

func main() {
	s1 := "abcde"
	s2 := "cdeab"

	fmt.Println(isCircular(s1, s2))
}

func isCircular(s1, s2 string) bool {
	for cnt, i := 0, 0; cnt < len(s1); cnt++ {
		newString := s1[i+1:] + string(s1[0])
		if newString == s2 {
			return true
		}
		s1 = newString
		cnt++
	}
	return false;
}