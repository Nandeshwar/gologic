package main

import (
	"fmt"
	"math"
)

func main() {
	r := isWordsOrderSame([]string{"ram", "shyam", "krishna"}, "srkhamin")
	fmt.Println(r)
}

func isWordsOrderSame(words []string, languageOrderStr string) bool {

	var languageOrder [26]int
	for i, v := range languageOrderStr {
		languageOrder[v-'a'] = i
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			iWord := words[i]
			jWord := words[j]

			l1 := len(words[i])
			l2 := len(words[j])

			min := int(math.Max(float64(l1), float64(l2)))

			for k := 0; k < min; k++ {
				if languageOrder[iWord[k]-'a'] < languageOrder[jWord[k]-'a'] {
					break
				} else if languageOrder[iWord[k]-'a'] > languageOrder[jWord[k]-'a'] {
					return false
				} else if k == min-1 && len(iWord) > len(jWord) {
					return false
				}
			}
		}
	}
	return true
}
