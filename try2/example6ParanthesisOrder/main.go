package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(isValidParenthesisOrder("{[]}"))
}

// valid: "{}"
// valid: "{[]}"
// valid: ""
// invalid: "{]"

func isValidParenthesisOrder(s string) bool {
	stack := list.New()

	if len(s) == 0 {
		return true
	}

	lookUp := map[string]string{"{": "}", "[": "]", "(": ")"}

	for _, v := range s {
		strV := string(v)
		if strV == "{" || strV == "[" || strV == "(" {
			stack.PushBack(strV)
		} else {
			element := stack.Remove(stack.Back())
			val := element.(string)

			if stack.Len() != 0 && lookUp[val] != strV {
				return false
			}
		}
	}
	return true
}
