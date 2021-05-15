package main

import (
	"fmt"
)

func main() {
	str1 := "Nandeshwar"
	str := []rune(str1)

	fp := 0;
	bp := len(str) - 1
	for fp < bp {
		/*
		tmp := str[fp];
		str[fp] = str[bp]
		str[bp] = tmp
		*/
		// above 3 lines or below 1 line
		str[fp], str[bp] = str[bp], str[fp]

		fp++
		bp--
	}
	fmt.Println(string(str))

	// just for fun
	fp = 0
	bp = 2;
	a := []int{10, 20, 30}
	a[fp], a[bp] = a[bp], a[fp]
	fmt.Println(a)

}

