package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "abca"
	fmt.Printf("Given string %#v is unique=%v", str, isStringUnique3(str))
}

// search using existing data structure - map
func isStringUnique1(str string) bool {
	m := map[string]string{}
	for i := 0; i < len(str); i++ {
		key := string(str[i])
		val := string(str[i])
		if val == m[key] {
			return false
		}
		m[key] = val
	}
	return true
}

// basic search without existing data structure
func isStringUnique2(str string) bool {
	var vArr []string

	exists := func(val string, vArr []string) bool {
		for i := 0; i < len(vArr); i++ {
			if val == vArr[i] {
				return true
			}
		}
		return false
	}

	for i := 0; i < len(str); i++ {
		val := string(str[i])
		if exists(val, vArr) {
			return false
		}
		vArr = append(vArr, val)
	}
	return true
}

// based on binary search
func isStringUnique3(str string) bool {

	var vArr []string

	exists := func(val string, vArr []string) bool {
		sort.Strings(vArr)

		begin := 0
		end := len(vArr)

		for begin < end {
			mid := (begin + end) / 2

			if val == vArr[mid] {
				return true
			}
			if val < vArr[mid] {
				end = mid
			} else {
				begin = mid + 1
			}
		}
		return false
	}

	for i := 0; i < len(str); i++ {
		val := string(str[i])
		if exists(val, vArr) {
			return false
		}

		vArr = append(vArr, val)
	}
	return true
}
