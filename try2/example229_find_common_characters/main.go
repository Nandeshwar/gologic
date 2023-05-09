package main

import "fmt"

func main() {
	//strList := []string{"bella", "label", "roller"} // output: ell
	strList := []string{"aab", "ac", "ba"} // output: a
	
	/*
	output: 
	  bash-3.2$ go run main.go
charFrequencies [2 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
charFrequencies [1 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
charFrequencies [1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
--------
minFrequency= [1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
result= a

	*/

	maxInt := int(^uint(0) >> 1)
	minFrequencies := make([]int, 26)
	for i := 0; i < 26; i++ {
		minFrequencies[i] = maxInt
	}

	for _, v := range strList {
		charFrequencies := make([]int, 26)

		for _, c := range v {
			charFrequencies[c-'a']++
		}

		// firstTime, all minFrequinces will either turn to 0 or >0 because min frequencies have either 0 or greater than 0
		
		// Always update minFrequencies,
		// if a letter is not present in charFrequencies, then it will turn to 0 in min frequency
		for i := 0; i < 26; i++ {
			minFrequencies[i] = min(minFrequencies[i], charFrequencies[i])
		}
		fmt.Println("charFrequencies", charFrequencies)
	}

	fmt.Println("--------")
	fmt.Println("minFrequency=", minFrequencies)

	var result []byte
	for i := 0; i < 26; i++ {

		for minFrequencies[i] > 0 {
			result = append(result, byte(i+'a'))
			minFrequencies[i]--
		}
	}
	fmt.Println("result=", string(result))

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
