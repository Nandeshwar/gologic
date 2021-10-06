package main

import (
	"fmt"
	"math"
)

func main() {
	//strSlice := []string{"ram", "shyam", "krishna", "radha"}
	strSlice := []string{"rama", "shyama", "krishnaa", "radhaa"}

	charCountArrMinVal := [26]int{}
	maxInt := int(^uint(0) >> 1)
	for i, _ := range charCountArrMinVal {
		charCountArrMinVal[i] = maxInt
	}
	for _, str := range strSlice {
		charArr := []rune(str)
		charCountArr := [26]int{}
		for _, c := range charArr {
			//charCountArr[c-97]++
			// or
			charCountArr[c-'a']++
		}

		for index, _ := range charCountArrMinVal {
			charCountArrMinVal[index] = int(math.Min(float64(charCountArr[index]), float64(charCountArrMinVal[index])))
		}
	}

	repeatedCharSlice := []rune{}
	for index, v := range charCountArrMinVal {
		for v > 0 {
			repeatedCharSlice = append(repeatedCharSlice, rune(index+97))
			v--
		}
	}

	fmt.Println(string(repeatedCharSlice))

}
