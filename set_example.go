package main

import "fmt"

func main() {

	a := UniqStr([]string {"10", "10", "20", "10"})
	fmt.Println(a)
}

func UniqStr(col []string) []string {
	m := map[string]struct{}{}
	for _, v := range col {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	list := make([]string, len(m))
	i := 0
	for  key := range m {
		fmt.Println(key)
		list[i] = key
		i++
	}
	return list
}