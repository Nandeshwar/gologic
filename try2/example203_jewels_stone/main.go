package main

import (
	"fmt"
	"strings"
)

func main() {
	jewels := "aa"
	stones := "baake"
	
	// output: 2

	// Algorithm1
	cnt := 0
	for _, c := range stones {
		if strings.Index(jewels, string(c)) > -1 {
			cnt++
		}
	}
	fmt.Println("cnt1=", cnt)

	// Algorithm2: using space
	cnt = 0
	m := map[byte]struct{}{}
	for _, c := range stones {
		m[byte(c)] = struct{}{}
	}

	for _, v := range jewels {
		_, ok := m[byte(v)]
		if ok {
			cnt++
		}
	}
	fmt.Println("cnt2=", cnt)
}
