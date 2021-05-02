// https://www.youtube.com/watch?v=UPdSViixmDs&list=PLU_sdQYzUj2keVENTP0a5rdykRSgg9Wp-&index=12
package main

import "fmt"

func main() {
	fmt.Println(palindrome(121))
}

func palindrome(num int) bool {
	if num == 0 || num %10 == 0 || num < 0 {
		return false
	}

	reverseNum := 0

	for num > reverseNum {
		r := num % 10
		num = num / 10
		reverseNum = reverseNum * 10 + r
	}

	if reverseNum == num || reverseNum / 10 == num {
		return true
	}

	return false
}