package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(findAllAnagram([]string{"ram", "sita", "Radha", "mar", "ahdaR", "shyam"}))
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
