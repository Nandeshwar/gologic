package main

import "fmt"

func main() {

	s := "++++"

	var result []string

	for i := 0; i < len(s)-1; i++ {
		if string(s[i]) == "+" && string(s[i+1]) == "+" {
			result = append(result, string(s[:i])+"--"+string(s[i+2:]))
		}
	}
	fmt.Println(result)
}

/*
output: [--++ +--+ ++--]
*/
