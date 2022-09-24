package main

import (
	"fmt"
	"strconv"
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

	key := string(keys[0])
	keyInt, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("error=", err.Error())
	}

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
