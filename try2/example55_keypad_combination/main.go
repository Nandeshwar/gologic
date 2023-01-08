package main

import (
	"fmt"
)

var code = []string{"abc", "cde", "efg", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz"}

func main() {
	// output: [psv psw psx ptv ptw ptx puv puw pux qsv qsw qsx qtv qtw qtx quv quw qux rsv rsw rsx rtv rtw rtx ruv ruw rux]
	fmt.Println(getKpc("678"))
}

func getKpc(keys string) []string {
	if len(keys) == 0 {
		return []string{""}
	}

	keyInt := int(keys[0] - '0')

	remaningKeys := string(keys[1:])
	remainingResult := getKpc(remaningKeys)

	resultArray := []string{}

	for _, v := range code[keyInt] {
		for _, r := range remainingResult {
			s := string(v) + string(r)
			resultArray = append(resultArray, s)
		}
	}
	return resultArray
}
