package main

import "fmt"

func main() {
	input := []string{"rama", "ram", "radha"}

	var result string

	var minWordLen = int(^uint(0) >> 1)

	for _, v := range input {
		minWordLen = min(len(v), minWordLen)
	}

	fmt.Println("minWordLen=", minWordLen)
outer:
	for i := 0; i < minWordLen; i++ {
		for j := 1; j < len(input); j++ {
			if string(input[0][i]) != string(input[j][i]) {
				break outer
			}
		}

		result += string(input[0][i])
	}

	fmt.Println("longest common prefix=", result)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
