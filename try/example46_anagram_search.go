package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	anagramsMap := map[string][]string{}
	anagramsStrList := []string{"Test", "estT", "TTes", "abc", "Cba", "Ram", "Sita", "Radha", "Dhara", "mar"}

	for _, str := range anagramsStrList {
		sortedWord := sortStr(str)
		vals, ok := anagramsMap[sortedWord]
		if ok {
			anagramsMap[sortStr(str)] = append(vals, str)
			continue
		} 
		anagramsMap[sortStr(str)] = []string{str}
	}

	for k, v := range anagramsMap {
		fmt.Printf("\n%v:%v", k, v)
	}
}

func sortStr(str string) string {
	str = strings.ToLower(str)
	letterList := strings.Split(str, "");
	sort.Strings(letterList)
	return strings.Join(letterList, "")
}