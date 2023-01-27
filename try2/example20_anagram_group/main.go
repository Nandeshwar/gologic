package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(findAllAnagram([]string{"ram", "sita", "Radha", "mar", "ahdaR", "shyam"}))
	fmt.Println(findAllAnagram2([]string{"ram", "sita", "Radha", "mar", "ahdaR", "shyam"}))
}

func findAllAnagram(strArr []string) []string {
	m := map[string][]string{}
	for _, str := range strArr {
		sortedStr := sortStr(str)
		v, ok := m[sortedStr]
		if ok {
			v = append(v, str)
			m[sortedStr] = v
		} else {
			m[sortedStr] = []string{str}
		}
	}
	var groupedAnagramList []string
	for _, v := range m {
		groupedAnagramList = append(groupedAnagramList, v...)
	}
	return groupedAnagramList

}

func sortStr(s string) string {
	strArr := strings.Split(s, "")
	sort.Strings(strArr)
	sortedStr := strings.Join(strArr, "")
	return sortedStr
}

func findAllAnagram2(strArr []string) []string {
	var result []string
	m := make(map[string][]string)
	for _, str := range strArr {
		a := [26]int{}
		
		for _, c := range str {
			a[unicode.ToLower(c) - 'a']++
		}
		
		var key string
		for _, num := range a {
			key += fmt.Sprintf("%d", num)
		}
		
		mv, ok := m[key] 
		if ok {
			mv = append(mv, str)
			m[key] = mv
		} else {
			mv = []string{str}
			m[key] = mv
		}
	}
	
	for _, v := range m {
		result = append(result, v...)
	}
	return result
}
