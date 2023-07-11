package main

import (
	"fmt"
)

func main() {
	num := 102
	roman := intToRoman(num)

	fmt.Printf("\nnum=%d, roman=%s\n", num, roman)

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
