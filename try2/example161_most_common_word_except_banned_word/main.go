package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}

	// expectation: ball this repeats twice
	commonWord := mostCommonWord(paragraph, banned)
	fmt.Println("Common word=", commonWord)

}

func mostCommonWord(paragraph string, banned []string) string {
	var commonWord string
	r, err := regexp.Compile("[a-z|A-Z]+")
	if err != nil {
		fmt.Println("error find word using regular expression")
		return ""
	}
	allStrings := r.FindAllString(paragraph, -1)

	// bannedWordMap := make(map[string]struct{})
	// or
	bannedWordMap := map[string]struct{}{} // struct{} type for empty struct last {} for map data
	for _, v := range banned {
		bannedWordMap[strings.ToLower(v)] = struct{}{} // struct{}{} is empty struct
	}

	wordsMap := make(map[string]int)
	for _, v := range allStrings {
		vLower := strings.ToLower(v)

		_, ok := bannedWordMap[vLower]
		if ok {
			continue
		}

		wordsCount, ok := wordsMap[vLower]
		if ok {
			wordsCount++
			wordsMap[vLower] = wordsCount
		} else {
			wordsMap[vLower] = 1
		}
	}

	fmt.Println("all strings=", allStrings)

	fmt.Println(commonWord)
	count := 0
	for k, v := range wordsMap {
		if v > count {
			count = v
			commonWord = k
		}
	}

	return commonWord
}
