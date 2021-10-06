package main

import (
	"fmt"
)

func main() {
	morseCode := []string {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",
	".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

	str := "nan"
	charArr := []rune(str)

	fmt.Println(morseCode)
	for _, c := range charArr {
		fmt.Println(morseCode[c-97])
	}
}