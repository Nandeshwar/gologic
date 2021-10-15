package main

func findSubstring(s, substring string) bool {
	substringCnt := len(substring)
	matchCount := 0

	sArr := []rune(s)
	substringArr := []rune(substring)

	for i, j := 0, 0; i < len(s); i++ {
		if sArr[i] == substringArr[j] {
			j++
			matchCount++

			if matchCount == substringCnt {
				return true
			}
		} else {
			j = 0
			matchCount = 0
		}
	}
	return false
}
