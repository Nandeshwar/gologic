package main

import "fmt"

func main() {
	fmt.Println(isAnagrams("eye", "yee"))

}

func isAnagrams(str1, str2 string) bool {
	str1Len := len(str1)
	str2Len := len(str2)

	if str1Len != str2Len {
		return false
	}

	for i := 0; i < str1Len; i++ {
		str1Count := 0
		str2Count := 0
		for j := 0; j < str1Len; j++ {
			if str1[i] == str1[j] {
				str1Count++
			}
			if str1[i] == str2[j] {
				str2Count++
			}
		}

		if str1Count != str2Count {
			return false
		}
	}
	return true
}
