package main

import (
	"container/list"
	"fmt"
)

func main() {
	str1 := "ab#c"
	str2 := "abb##c"

	// space: o(n)
	fmt.Println(isBothStrSame(str1, str2))
	// space: o(1)
	fmt.Println(isBothStrSame2(str1, str2))
}

// o(n) time, space o(n)
func isBothStrSame(str1, str2 string) bool {
	s1 := list.New()
	s2 := list.New()

	for _, v := range str1 {
		strV := string(v)
		switch strV {
		case "#":
			if s1.Len() != 0 {
				s1.Remove(s1.Back())
			}
		default:
			s1.PushBack(strV)
		}
	}

	for _, v := range str2 {
		strV := string(v)
		switch string(v) {
		case "#":
			if s2.Len() != 0 {
				s2.Remove(s2.Back())
			}
		default:
			s2.PushBack(strV)
		}
	}

	if s1.Len() != s1.Len() {
		return false
	}

	for i := 0; i < s1.Len(); i++ {
		element1 := s1.Remove(s1.Back())
		element2 := s2.Remove(s2.Back())

		item1 := element1.(string)
		item2 := element2.(string)

		if item1 != item2 {
			return false
		}
	}
	return true
}

// o(n) time, space o(n)
func isBothStrSame2(str1, str2 string) bool {
	i := len(str1) - 1
	j := len(str2) - 1

	for i >= 0 || j >= 0 {
		cnt := 0
		for i >= 0 && (cnt > 0 || string(str1[i]) == "#") {
			if string(str1[i]) == "#" {
				cnt++
			} else {
				cnt--
			}
			i--
		}

		for j >= 0 && (cnt > 0 || string(str2[j]) == "#") {

			if string(str2[j]) == "#" {
				cnt++
			} else {
				cnt--
			}
			j--
		}

		if i >= 0 && j >= 0 {

			if str1[i] != str2[j] {
				return false
			} else {
				i--
				j--
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
	}

	return true
}
