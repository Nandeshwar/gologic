package main

import (
	"fmt"
)

// validate paranthesis. * can be any of paranthesis
/*
	Given a string s containing only three types of characters: '(', ')' and '*', return true if s is valid.

The following rules define a valid string:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string "".


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "(*)"
Output: true
Example 3:

Input: s = "(*))"
Output: true
*/

func main() {
	s := "(*)"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	cnt := 0

	for _, v := range s {
		if v == ')' {
			cnt--
		} else {
			cnt++
		}

		if cnt < 0 {
			return false
		}
	}

	if cnt == 0 {
		return true
	}

	fmt.Println("2nd portion")
	cnt = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			cnt--
		} else {
			cnt++
		}

		if cnt < 0 {
			return false
		}
	}

	return true

}
