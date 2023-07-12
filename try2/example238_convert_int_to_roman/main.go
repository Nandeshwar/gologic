package main

import (
	"fmt"
)

func main() {
	num := 102 // CII
	roman := intToRoman(num)

	fmt.Printf("\nnum=%d, roman=%s\n", num, roman)

	romanStr := "CII"    // 102
	romanStr = "MCMXCIV" // 1994
	intReslult := convertRomanToInt(romanStr)

	fmt.Println("\nromanStr=%s, num=%d", romanStr, intReslult)

}

func intToRoman(num int) string {
	var romanStr string

	intList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for ind, v := range intList {
		for num >= v {
			romanStr += roman[ind]
			num = num - v
		}
	}
	return romanStr
}

func convertRomanToInt(roman string) int {
	result := 0

	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := len(roman) - 1; i >= 0; i-- {
		s := string(roman[i])

		v := m[s]
		if v >= result {
			result += v
		} else {
			result -= v
		}
	}
	return result
}
