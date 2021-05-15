package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(isStrSame2("ab#c", "ad#c"))
	fmt.Println(isStrSame2("abb#c", "adb#c"))

}

func isStrSame2(s1, s2 string) bool {
	stack1 := list.New()
	stack2 := list.New()

	for _, c := range s1 {
		if c != '#' {
			stack1.PushBack(c)
		} else {
			stack1.Remove(stack1.Back())
		}
	}

	for _, c := range s2 {
		if c != '#' {
			stack2.PushBack(c)
		} else {
			stack2.Remove(stack2.Back())
		}
	}

	for stack1.Len() != 0 && stack2.Len() != 0 {
		ch1 := stack1.Remove(stack1.Back())
		ch2 := stack2.Remove(stack2.Back())

		if ch1 != ch2 {
			return false
		}
	}

	return true
}