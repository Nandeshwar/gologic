package main

import (
	"fmt"
)

func main() {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..",
		".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.",
		"...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	inputWords := []string{"gin", "zen", "gig", "msg"}
	
	// expected count = 2

	cnt := findUniqueMorseCodeCount(inputWords, morseCodes)
	fmt.Println("cnt=", cnt)
}

func findUniqueMorseCodeCount(inputWords, morseCodes []string) int {
	m := make(map[string]struct{})
	for _, str := range inputWords {
		var morseCodeStr string
		for _, v := range str {
			morseCodeStr += morseCodes[v-'a']
		}
		m[morseCodeStr] = struct{}{}
	}
	return len(m)
}
