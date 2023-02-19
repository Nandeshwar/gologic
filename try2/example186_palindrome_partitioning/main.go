package main

import (
	"fmt"
)

func main() {
	str := "aab"
	/*
		output:
		result= [[a a b] [aa b]]
		  [a a b] : partitioned 3 places and all are palindrome
		  [aa b]: paritined in 2 places are all are palindrome
	*/

	var result [][]string
	var path []string
	allPalindrome(str, &result, path, 0)
	fmt.Println("result=", result)
}

func allPalindrome(str string, result *[][]string, path []string, idx int) {
	if idx == len(str) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	// loop is for paritioning. 
	// 1. first iteration: partition at 0 index, check string 0 to 0 + 1 is palindrome
	// 2. 2nd iteration:   partition at 0 index, check string 0 to 0 + 2 is palindrome
	// 3. 3rd iteration:   partition at 0 index, check string 0 to 0 + 3 is palindrome
	//   At every iteration there is recursive call
	for i := idx; i < len(str); i++ {
		if isPalindrome(str[idx : i+1]) {
			path = append(path, str[idx:i+1])
			allPalindrome(str, result, path, i+1)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(str string) bool {
	i := 0
	j := len(str) - 1

	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
